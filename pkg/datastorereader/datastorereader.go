package datastorereader

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "strings"
)

const (
    existsSubpath = "/exists"
    getAllSubpath = "/get_all"
    filterSubpath = "/filter"
)

// Conn holds a connection to the datastore service (reader and writer).
type Conn struct {
    readerURL *url.URL
}

// New returns a new connection to the datastore.
func New(readerURL *url.URL) *Conn {
    d := new(Conn)
    d.readerURL = readerURL
    return d
}

// Exists does check if a collection object with given id exists.
func (d *Conn) GetAll(ctx context.Context, collection string) (bool, error) {
    reqBody := fmt.Sprintf(
        `{
            "collection": "%s"
        }`,
        collection
    )
    addr := d.readerURL.String() + getAllSubpath

    req, err := http.NewRequestWithContext(ctx, "POST", addr, strings.NewReader(reqBody))
    if err != nil {
        return false, fmt.Errorf("creating request to datastore: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return false, fmt.Errorf("sending request to datastore at %s: %w", addr, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            body = []byte("[can not read body]")
        }
        return false, fmt.Errorf("got response `%s`: %s", resp.Status, body)
    }

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return false, fmt.Errorf("reading response body: %w", err)
    }

    var respData struct {
        Data json.RawMessage `json:""`
    }
    if err := json.Unmarshal(respBody, &respData); err != nil {
        return false, fmt.Errorf("decoding response body `%s`: %w", respBody, err)
    }
    return respData.Data, nil
}

// Exists does check if a collection object with given id exists.
// TODO: filter should probably an interface or smth, @normanjaeckel please tell me where to define it
func (d *Conn) Exists(ctx context.Context, collection string, filter []string) (bool, error) {
    reqBody := fmt.Sprintf(
        `{
            "collection": "%s",
            "filter": {
                "field": "%s",
                "value": %s,
                "operator": "="
            }
        }`,
        collection, filter[0], filter[1]
    )
    addr := d.readerURL.String() + existsSubpath

    req, err := http.NewRequestWithContext(ctx, "POST", addr, strings.NewReader(reqBody))
    if err != nil {
        return false, fmt.Errorf("creating request to datastore: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return false, fmt.Errorf("sending request to datastore at %s: %w", addr, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            body = []byte("[can not read body]")
        }
        return false, fmt.Errorf("got response `%s`: %s", resp.Status, body)
    }

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return false, fmt.Errorf("reading response body: %w", err)
    }

    var respData struct {
        Exists bool `json:"exists"`
    }
    if err := json.Unmarshal(respBody, &respData); err != nil {
        return false, fmt.Errorf("decoding response body `%s`: %w", respBody, err)
    }
    return respData.Exists, nil
}


// Exists does check if a collection object with given id exists.
// TODO: filter should probably an interface or smth, @normanjaeckel please tell me where to define it
func (d *Conn) Filter(ctx context.Context, collection string, filter []string, fields []string) (bool, error) {

//> reader/filter { "collection": "meeting", "filter": { "field": "first_name", "value": "Abdullah", "operator": "=" }, "mapped_fields": ["email", "last_name"] }

    fieldsStr := ""
    if len(fields) > 0 {
        fieldsStr = ", \"fields\": [" +
            "\"" + strings.Join(fields, "\", \"") + "\"" +
            "]"
    }

    reqBody := fmt.Sprintf(
        `{
            "collection": "%s",
            "filter": {
                "field": "%s",
                "value": %s,
                "operator": "="
            }%s
        }`,
        collection, filter[0], filter[1], fieldsStr
    )
    addr := d.readerURL.String() + filterSubpath

    req, err := http.NewRequestWithContext(ctx, "POST", addr, strings.NewReader(reqBody))
    if err != nil {
        return false, fmt.Errorf("creating request to datastore: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return false, fmt.Errorf("sending request to datastore at %s: %w", addr, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            body = []byte("[can not read body]")
        }
        return false, fmt.Errorf("got response `%s`: %s", resp.Status, body)
    }

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return false, fmt.Errorf("reading response body: %w", err)
    }


//'{ "collection": "user", "filter": { "field": "first_name", "value": "Abdullah", "operator": "=" }, "mapped_fields": ["last_name", "email"] }'
//'{"data":{"55":{"last_name":"\u015eevik","email":"a.sevik@me.com"}},"position":1342}'



    var respData struct {
        Exists bool `json:"exists"`
    }
    if err := json.Unmarshal(respBody, &respData); err != nil {
        return false, fmt.Errorf("decoding response body `%s`: %w", respBody, err)
    }
    return respData.Exists, nil
}



// Create sends multiple create events to the datastore.
func (d *Conn) Create(ctx context.Context, creatables map[string]map[string]json.RawMessage, migrationIndex int) error {
    var events []json.RawMessage
    for fqid, fields := range creatables {
        event := struct {
            Type   string                     `json:"type"`
            Fqid   string                     `json:"fqid"`
            Fields map[string]json.RawMessage `json:"fields"`
        }{
            Type:   "create",
            Fqid:   fqid,
            Fields: fields,
        }
        encodedEvent, err := json.Marshal(event)
        if err != nil {
            return fmt.Errorf("marshalling create event: %w", err)
        }
        events = append(events, encodedEvent)
    }

    reqBody := struct {
        UserID         int               `json:"user_id"`
        Information    map[string]string `json:"information"`
        LockedFields   map[string]string `json:"locked_fields"`
        Events         []json.RawMessage `json:"events"`
        MigrationIndex int               `json:"migration_index"`
    }{
        0,
        map[string]string{},
        map[string]string{},
        events,
        migrationIndex,
    }
    encodedReqBody, err := json.Marshal(reqBody)
    if err != nil {
        return fmt.Errorf("marshalling write request: %w", err)

    }

    if err := sendWriteRequest(ctx, d.writerURL, string(encodedReqBody)); err != nil {
        return fmt.Errorf("sending write request to datastore: %w", err)
    }

    return nil
}

// Set sends an update event to the datastore to set a FQField. The value has to be JSON.
func (d *Conn) Set(ctx context.Context, fqfield string, value json.RawMessage) error {
    parts := strings.Split(fqfield, "/")
    if len(parts) != 3 {
        return fmt.Errorf("invalid FQField %s, expected two `/`", fqfield)
    }

    reqBody := fmt.Sprintf(
        `{
            "user_id": 0,
            "information": {},
            "locked_fields":{}, "events":[
                {"type":"update","fqid":"%s/%s","fields":{"%s":%s}}
            ]
        }`,
        parts[0], parts[1], parts[2], value,
    )

    if err := sendWriteRequest(ctx, d.writerURL, reqBody); err != nil {
        return fmt.Errorf("sending write request to datastore: %w", err)
    }

    return nil
}

// sendWriteRequest sends the give request body to the datastore.
func sendWriteRequest(ctx context.Context, writerURL *url.URL, reqBody string) error {
    addr := writerURL.String() + writeSubpath

    req, err := http.NewRequestWithContext(ctx, "POST", addr, strings.NewReader(reqBody))
    if err != nil {
        return fmt.Errorf("creating request to datastore: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return fmt.Errorf("sending request to datastore at %s: %w", addr, err)
    }

    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            body = []byte("[can not read body]")
        }
        return fmt.Errorf("got response `%s`: %s", resp.Status, body)
    }

    return nil
}
