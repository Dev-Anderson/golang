package models

import (
	"database/sql"

	_struct "api-firebird/struct"
)

type ModelGetData struct {
	DB *sql.DB
}

func (model ModelGetData) GetNotaFiscal() (getStruct []_struct.StructData, err error) {
	row, err := model.DB.Query("SELECT NFE_TBNFESIDE.NUMERO, NFE_TBNFESIDE.SERIE, NFE_TBNFES.ID FROM NFE_TBNFES INNER JOIN NFE_TBNFESIDE ON NFE_TBNFES.GID = NFE_TBNFESIDE.GID WHERE NFE_TBNFES.NUMEROPROTOCOLO IS NOT NULL AND NFE_TBNFES.GID = 1")
	if err != nil {
		return nil, err
	} else {
		var _isiStruct []_struct.StructData
		var data _struct.StructData
		for row.Next() {
			err2 := row.Scan(
				&data.Numero,
				&data.Serie,
				&data.Id,
			)
			if err2 != nil {
				return nil, err2
			} else {
				_data := _struct.StructData{
					Numero: data.Numero,
					Serie:  data.Serie,
					Id:     data.Id,
				}
				_isiStruct = append(_isiStruct, _data)
			}
		}
		return _isiStruct, nil
	}
}
