package models

import (
	"database/sql"

	_struct "api-firebird/struct"
)

type ModelGetData struct {
	DB *sql.DB
}

func (model ModelGetData) GetNotasAprovadas() (getStruct []_struct.StructData, err error) {
	row, err := model.DB.Query("SELECT    NFES.GID,    IDE.NUMERO,    IDE.SERIE,    NFES.ID,    NFES.NUMEROPROTOCOLO,    NFES.DIGVALUE,    IDE.DATAEMISSAO,    IDE.HORAEMISSAO FROM NFE_TBNFES NFES INNER JOIN NFE_TBNFESIDE IDE ON NFES.GID = IDE.GID WHERE NFES.NUMEROPROTOCOLO IS NOT NULL AND NFES.DIGVALUE IS NOT NULL")
	if err != nil {
		return nil, err
	} else {
		var _isiStruct []_struct.StructData
		var data _struct.StructData
		for row.Next() {
			err2 := row.Scan(
				&data.Gid,
				&data.Numero,
				&data.Serie,
				&data.Id,
				&data.NumeroProtocolo,
				&data.DigValue,
				&data.DataEmissao,
				&data.HoraEmissao,
			)
			if err2 != nil {
				return nil, err2
			} else {
				_data := _struct.StructData{
					Gid:             data.Gid,
					Numero:          data.Numero,
					Serie:           data.Serie,
					Id:              data.Id,
					NumeroProtocolo: data.NumeroProtocolo,
					DigValue:        data.DigValue,
					DataEmissao:     data.DataEmissao,
					HoraEmissao:     data.HoraEmissao,
				}
				_isiStruct = append(_isiStruct, _data)
			}
		}
		return _isiStruct, nil
	}
}

func (model ModelGetData) GetNotasCanceladas() (getStruct []_struct.StructDataCanceladas, err error) {
	row, err := model.DB.Query("SELECT    NFES.GID,    IDE.NUMERO,    IDE.SERIE,    NFES.ID,    IDE.DATAEMISSAO,    IDE.HORAEMISSAO,    CAN.CANCELADOEM FROM NFE_TBNFES NFES INNER JOIN NFE_TBNFESIDE IDE ON NFES.GID = IDE.GID INNER JOIN NFE_TBCANCELAMENTONFE CAN ON NFES.GID = CAN.GID WHERE CAN.CANCELADOEM IS NOT NULL")
	if err != nil {
		return nil, err
	} else {
		var _isiStruct []_struct.StructDataCanceladas
		var data _struct.StructDataCanceladas
		for row.Next() {
			err2 := row.Scan(
				&data.Gid,
				&data.Numero,
				&data.Serie,
				&data.Id,
				&data.DataEmissao,
				&data.HoraEmissao,
				&data.CanceladoEm,
			)
			if err2 != nil {
				return nil, err2
			} else {
				_data := _struct.StructDataCanceladas{
					Gid:         data.Gid,
					Numero:      data.Numero,
					Serie:       data.Serie,
					Id:          data.Id,
					DataEmissao: data.DataEmissao,
					HoraEmissao: data.HoraEmissao,
					CanceladoEm: data.CanceladoEm,
				}
				_isiStruct = append(_isiStruct, _data)
			}
		}
		return _isiStruct, nil
	}
}
