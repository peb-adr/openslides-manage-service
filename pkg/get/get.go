package get

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "strings"

    "github.com/peb-adr/openslides-manage-service/pkg/connection"
    "github.com/peb-adr/openslides-manage-service/proto"
    "github.com/ghodss/yaml"
    "github.com/spf13/cobra"
    "google.golang.org/grpc"
    "google.golang.org/grpc/status"
)

const (
    // GetHelp contains the short help text for the command.
    GetHelp = "Get a property of a model from the datastore"

    // GetHelpExtra contains the long help text for the command without
    // the headline.
    GetHelpExtra = `Get a property of a model from the datastore`
)

// Cmd returns the create-user subcommand.
func Cmd(cmd *cobra.Command, cfg connection.Params) *cobra.Command {
    cmd.Use = "get"
    cmd.Short = GetHelp
    cmd.Long = GetHelp + "\n\n" + GetHelpExtra
    cmd.Args = cobra.ExactArgs(1)

    propHelpText := "Prop name to look up"

    cmd.RunE = func(cmd *cobra.Command, args []string) error {
        ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout())
        defer cancel()

        cl, close, err := connection.Dial(ctx, cfg.Addr(), cfg.PasswordFile(), !cfg.NoSSL())
        if err != nil {
            return fmt.Errorf("connecting to gRPC server: %w", err)
        }
        defer close()

        if err := Run(ctx, cl); err != nil {
            return fmt.Errorf("getting property: %w", err)
        }
        return nil
    }
    return cmd
}

// Client

type gRPCClient interface {
    Get(ctx context.Context, in *proto.GetRequest, opts ...grpc.CallOption) (*proto.GetResponse, error)
}

// Run calls respective procedure to set password of the given user.
func Run(ctx context.Context, gc gRPCClient, prop string) error {
    in := &proto.GetRequest{}
    in.Prop = prop

    resp, err := gc.Get(ctx, in)
    if err != nil {
        s, _ := status.FromError(err) // The ok value does not matter here.
        return fmt.Errorf("calling manage service: %s", s.Message())
    }
    fmt.Printf("%s\n", resp.Value)

    return nil
}

// Server

type datastore interface {
    GetAll(ctx context.Context, collection string) (bool, error)
    Exists(ctx context.Context, collection string, filter []string) (bool, error)
    Filter(ctx context.Context, collection string, filter []string, fields []string) (bool, error)
}

// CreateUser creates the given user.
// This function is the server side entrypoint for this package.
func CreateUser(ctx context.Context, in *proto.CreateUserRequest, a action) (*proto.CreateUserResponse, error) {
    name := "user.create"
    payload := []*proto.CreateUserRequest{in}
    data, err := json.Marshal(payload)
    if err != nil {
        return nil, fmt.Errorf("marshalling action data: %w", err)
    }
    result, err := a.Single(ctx, name, transform(data))
    if err != nil {
        return nil, fmt.Errorf("requesting backend action %q: %w", name, err)
    }

    var ids []struct {
        ID int `json:"id"`
    }
    if err := json.Unmarshal(result, &ids); err != nil {
        return nil, fmt.Errorf("unmarshalling action result %q: %w", string(result), err)
    }
    if len(ids) != 1 {
        return nil, fmt.Errorf("wrong lenght of action result, expected 1 item, got %d", len(ids))
    }
    return &proto.CreateUserResponse{UserID: int64(ids[0].ID)}, nil
}

// transform changes some JSON keys so we can use OpenSlides' template fields.
func transform(b []byte) []byte {
    fields := map[string]string{
        "committee__management_level": "committee_$_management_level",
        "group__ids":                  "group_$_ids",
    }
    s := string(b)
    for old, new := range fields {
        s = strings.ReplaceAll(s,
            fmt.Sprintf(`"%s":`, old),
            fmt.Sprintf(`"%s":`, new),
        )

    }
    return []byte(s)
}
