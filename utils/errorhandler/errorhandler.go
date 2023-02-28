package errorhandler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mattn/go-sqlite3"
)

var (
    ErrDup      = errors.New("このデータは既に登録済みです")
    ErrNoRecord = errors.New("該当のデータが見つかりませんでした")
	ErrPaswordUnmatch = errors.New("パスワードが一致しません")
)

func WrapDBError(err error) error {
    var sqliteErr sqlite3.Error
    if errors.As(err, &sqliteErr) {
        if errors.Is(sqliteErr.Code, sqlite3.ErrConstraint) {
        	return ErrDup
        }
    } else if errors.Is(err, sql.ErrNoRows) {
        return ErrNoRecord
    }
    return err
}

func MakeErrResponse(err error,w http.ResponseWriter,errCode int) {
	w.WriteHeader(errCode)
	resp := make(map[string]string)
	resp["message"] = err.Error()
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}