package models

import (
	fmt "fmt"
	time "time"

	git_querycap_com_tools_datatypes "git.querycap.com/tools/datatypes"
	github_com_go_courier_sqlx_v2 "github.com/go-courier/sqlx/v2"
	github_com_go_courier_sqlx_v2_builder "github.com/go-courier/sqlx/v2/builder"
	github_com_go_courier_sqlx_v2_datatypes "github.com/go-courier/sqlx/v2/datatypes"
)

func (Firmware) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Firmware) UniqueIndexUIFirmwareID() string {
	return "ui_firmware_id"
}

func (Firmware) UniqueIndexUIResID() string {
	return "ui_res_id"
}

func (Firmware) UniqueIndexUIVersion() string {
	return "ui_version"
}

func (Firmware) UniqueIndexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"ui_firmware_id": []string{
			"FirmwareID",
		},
		"ui_res_id": []string{
			"ResID",
		},
		"ui_version": []string{
			"Name",
			"Major",
			"Minor",
			"Patch",
			"Identifier",
			"ResID",
		},
	}
}

func (Firmware) Comments() map[string]string {
	return map[string]string{
		"CreatedAt":  "创建时间",
		"Desc":       "固件描述",
		"FirmwareID": "固件ID",
		"Identifier": "修饰符",
		"Major":      "主版本号",
		"Minor":      "次版本号",
		"Name":       "固件名称",
		"Patch":      "修订号",
		"ReleaseAt":  "发布时间",
		"ResID":      "资源ID",
		"UpdatedAt":  "更新时间",
	}
}

var FirmwareTable *github_com_go_courier_sqlx_v2_builder.Table

func init() {
	FirmwareTable = DB.Register(&Firmware{})
}

type FirmwareIterator struct {
}

func (FirmwareIterator) New() interface{} {
	return &Firmware{}
}

func (FirmwareIterator) Resolve(v interface{}) *Firmware {
	return v.(*Firmware)
}

func (Firmware) TableName() string {
	return "t_firmware"
}

func (Firmware) ColDescriptions() map[string][]string {
	return map[string][]string{
		"CreatedAt": []string{
			"创建时间",
		},
		"Desc": []string{
			"固件描述",
		},
		"FirmwareID": []string{
			"固件ID",
		},
		"Identifier": []string{
			"修饰符",
		},
		"Major": []string{
			"主版本号",
		},
		"Minor": []string{
			"次版本号",
		},
		"Name": []string{
			"固件名称",
		},
		"Patch": []string{
			"修订号",
		},
		"ReleaseAt": []string{
			"发布时间",
		},
		"ResID": []string{
			"资源ID",
		},
		"UpdatedAt": []string{
			"更新时间",
		},
	}
}

func (Firmware) FieldKeyID() string {
	return "ID"
}

func (m *Firmware) FieldID() *github_com_go_courier_sqlx_v2_builder.Column {
	return FirmwareTable.F(m.FieldKeyID())
}

func (Firmware) FieldKeyFirmwareID() string {
	return "FirmwareID"
}

func (m *Firmware) FieldFirmwareID() *github_com_go_courier_sqlx_v2_builder.Column {
	return FirmwareTable.F(m.FieldKeyFirmwareID())
}

func (Firmware) FieldKeyResID() string {
	return "ResID"
}

func (m *Firmware) FieldResID() *github_com_go_courier_sqlx_v2_builder.Column {
	return FirmwareTable.F(m.FieldKeyResID())
}

func (Firmware) FieldKeyMajor() string {
	return "Major"
}

func (m *Firmware) FieldMajor() *github_com_go_courier_sqlx_v2_builder.Column {
	return FirmwareTable.F(m.FieldKeyMajor())
}

func (Firmware) FieldKeyMinor() string {
	return "Minor"
}

func (m *Firmware) FieldMinor() *github_com_go_courier_sqlx_v2_builder.Column {
	return FirmwareTable.F(m.FieldKeyMinor())
}

func (Firmware) FieldKeyPatch() string {
	return "Patch"
}

func (m *Firmware) FieldPatch() *github_com_go_courier_sqlx_v2_builder.Column {
	return FirmwareTable.F(m.FieldKeyPatch())
}

func (Firmware) FieldKeyIdentifier() string {
	return "Identifier"
}

func (m *Firmware) FieldIdentifier() *github_com_go_courier_sqlx_v2_builder.Column {
	return FirmwareTable.F(m.FieldKeyIdentifier())
}

func (Firmware) FieldKeyName() string {
	return "Name"
}

func (m *Firmware) FieldName() *github_com_go_courier_sqlx_v2_builder.Column {
	return FirmwareTable.F(m.FieldKeyName())
}

func (Firmware) FieldKeyDesc() string {
	return "Desc"
}

func (m *Firmware) FieldDesc() *github_com_go_courier_sqlx_v2_builder.Column {
	return FirmwareTable.F(m.FieldKeyDesc())
}

func (Firmware) FieldKeyReleaseAt() string {
	return "ReleaseAt"
}

func (m *Firmware) FieldReleaseAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return FirmwareTable.F(m.FieldKeyReleaseAt())
}

func (Firmware) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *Firmware) FieldCreatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return FirmwareTable.F(m.FieldKeyCreatedAt())
}

func (Firmware) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Firmware) FieldUpdatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return FirmwareTable.F(m.FieldKeyUpdatedAt())
}

func (Firmware) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *Firmware) IndexFieldNames() []string {
	return []string{
		"FirmwareID",
		"ID",
		"Identifier",
		"Major",
		"Minor",
		"Name",
		"Patch",
		"ResID",
	}
}

func (m *Firmware) ConditionByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) github_com_go_courier_sqlx_v2_builder.SqlCondition {
	table := db.T(m)
	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m)

	conditions := make([]github_com_go_courier_sqlx_v2_builder.SqlCondition, 0)

	for _, fieldName := range m.IndexFieldNames() {
		if v, exists := fieldValues[fieldName]; exists {
			conditions = append(conditions, table.F(fieldName).Eq(v))
			delete(fieldValues, fieldName)
		}
	}

	if len(conditions) == 0 {
		panic(fmt.Errorf("at least one of field for indexes has value"))
	}

	for fieldName, v := range fieldValues {
		conditions = append(conditions, table.F(fieldName).Eq(v))
	}

	condition := github_com_go_courier_sqlx_v2_builder.And(conditions...)

	return condition
}

func (m *Firmware) Create(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_go_courier_sqlx_v2.InsertToDB(db, m, nil))
	return err

}

func (m *Firmware) CreateOnDuplicateWithUpdateFields(db github_com_go_courier_sqlx_v2.DBExecutor, updateFields []string) error {

	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, updateFields...)

	delete(fieldValues, "ID")

	table := db.T(m)

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	fields := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		fields[field] = true
	}

	for _, fieldNames := range m.UniqueIndexes() {
		for _, field := range fieldNames {
			delete(fields, field)
		}
	}

	if len(fields) == 0 {
		panic(fmt.Errorf("no fields for updates"))
	}

	for field := range fieldValues {
		if !fields[field] {
			delete(fieldValues, field)
		}
	}

	additions := github_com_go_courier_sqlx_v2_builder.Additions{}

	switch db.Dialect().DriverName() {
	case "mysql":
		additions = append(additions, github_com_go_courier_sqlx_v2_builder.OnDuplicateKeyUpdate(table.AssignmentsByFieldValues(fieldValues)...))
	case "postgres":
		indexes := m.UniqueIndexes()
		fields := make([]string, 0)
		for _, fs := range indexes {
			fields = append(fields, fs...)
		}
		indexFields, _ := db.T(m).Fields(fields...)

		additions = append(additions,
			github_com_go_courier_sqlx_v2_builder.OnConflict(indexFields).
				DoUpdateSet(table.AssignmentsByFieldValues(fieldValues)...))
	}

	additions = append(additions, github_com_go_courier_sqlx_v2_builder.Comment("User.CreateOnDuplicateWithUpdateFields"))

	expr := github_com_go_courier_sqlx_v2_builder.Insert().Into(table, additions...).Values(cols, vals...)

	_, err := db.ExecExpr(expr)
	return err

}

func (m *Firmware) DeleteByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(m.ConditionByStruct(db)),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.DeleteByStruct"),
			),
	)

	return err
}

func (m *Firmware) FetchByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.FetchByID"),
			),
		m,
	)

	return err
}

func (m *Firmware) UpdateByIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.UpdateByIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByID(db)
	}

	return nil

}

func (m *Firmware) UpdateByIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *Firmware) FetchByIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Firmware) DeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.DeleteByID"),
			))

	return err
}

func (m *Firmware) FetchByFirmwareID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("FirmwareID").Eq(m.FirmwareID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.FetchByFirmwareID"),
			),
		m,
	)

	return err
}

func (m *Firmware) UpdateByFirmwareIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("FirmwareID").Eq(m.FirmwareID),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.UpdateByFirmwareIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByFirmwareID(db)
	}

	return nil

}

func (m *Firmware) UpdateByFirmwareIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByFirmwareIDWithMap(db, fieldValues)

}

func (m *Firmware) FetchByFirmwareIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("FirmwareID").Eq(m.FirmwareID),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.FetchByFirmwareIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Firmware) DeleteByFirmwareID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("FirmwareID").Eq(m.FirmwareID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.DeleteByFirmwareID"),
			))

	return err
}

func (m *Firmware) FetchByResID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ResID").Eq(m.ResID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.FetchByResID"),
			),
		m,
	)

	return err
}

func (m *Firmware) UpdateByResIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("ResID").Eq(m.ResID),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.UpdateByResIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByResID(db)
	}

	return nil

}

func (m *Firmware) UpdateByResIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByResIDWithMap(db, fieldValues)

}

func (m *Firmware) FetchByResIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ResID").Eq(m.ResID),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.FetchByResIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Firmware) DeleteByResID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ResID").Eq(m.ResID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.DeleteByResID"),
			))

	return err
}

func (m *Firmware) FetchByNameAndMajorAndMinorAndPatchAndIdentifierAndResID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Name").Eq(m.Name),
					table.F("Major").Eq(m.Major),
					table.F("Minor").Eq(m.Minor),
					table.F("Patch").Eq(m.Patch),
					table.F("Identifier").Eq(m.Identifier),
					table.F("ResID").Eq(m.ResID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.FetchByNameAndMajorAndMinorAndPatchAndIdentifierAndResID"),
			),
		m,
	)

	return err
}

func (m *Firmware) UpdateByNameAndMajorAndMinorAndPatchAndIdentifierAndResIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("Name").Eq(m.Name),
					table.F("Major").Eq(m.Major),
					table.F("Minor").Eq(m.Minor),
					table.F("Patch").Eq(m.Patch),
					table.F("Identifier").Eq(m.Identifier),
					table.F("ResID").Eq(m.ResID),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.UpdateByNameAndMajorAndMinorAndPatchAndIdentifierAndResIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByNameAndMajorAndMinorAndPatchAndIdentifierAndResID(db)
	}

	return nil

}

func (m *Firmware) UpdateByNameAndMajorAndMinorAndPatchAndIdentifierAndResIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByNameAndMajorAndMinorAndPatchAndIdentifierAndResIDWithMap(db, fieldValues)

}

func (m *Firmware) FetchByNameAndMajorAndMinorAndPatchAndIdentifierAndResIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Name").Eq(m.Name),
					table.F("Major").Eq(m.Major),
					table.F("Minor").Eq(m.Minor),
					table.F("Patch").Eq(m.Patch),
					table.F("Identifier").Eq(m.Identifier),
					table.F("ResID").Eq(m.ResID),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.FetchByNameAndMajorAndMinorAndPatchAndIdentifierAndResIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Firmware) DeleteByNameAndMajorAndMinorAndPatchAndIdentifierAndResID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Name").Eq(m.Name),
					table.F("Major").Eq(m.Major),
					table.F("Minor").Eq(m.Minor),
					table.F("Patch").Eq(m.Patch),
					table.F("Identifier").Eq(m.Identifier),
					table.F("ResID").Eq(m.ResID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Firmware.DeleteByNameAndMajorAndMinorAndPatchAndIdentifierAndResID"),
			))

	return err
}

func (m *Firmware) List(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) ([]Firmware, error) {

	list := make([]Firmware, 0)

	table := db.T(m)
	_ = table

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("Firmware.List"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(db.T(m), finalAdditions...),
		&list,
	)

	return list, err

}

func (m *Firmware) Count(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("Firmware.Count"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(
			github_com_go_courier_sqlx_v2_builder.Count(),
		).
			From(db.T(m), finalAdditions...),
		&count,
	)

	return count, err

}

func (m *Firmware) BatchFetchByFirmwareIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []git_querycap_com_tools_datatypes.SFID) ([]Firmware, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("FirmwareID").In(values)

	return m.List(db, condition)

}

func (m *Firmware) BatchFetchByIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]Firmware, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *Firmware) BatchFetchByIdentifierList(db github_com_go_courier_sqlx_v2.DBExecutor, values []string) ([]Firmware, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Identifier").In(values)

	return m.List(db, condition)

}

func (m *Firmware) BatchFetchByMajorList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]Firmware, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Major").In(values)

	return m.List(db, condition)

}

func (m *Firmware) BatchFetchByMinorList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]Firmware, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Minor").In(values)

	return m.List(db, condition)

}

func (m *Firmware) BatchFetchByNameList(db github_com_go_courier_sqlx_v2.DBExecutor, values []string) ([]Firmware, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Name").In(values)

	return m.List(db, condition)

}

func (m *Firmware) BatchFetchByPatchList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]Firmware, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Patch").In(values)

	return m.List(db, condition)

}

func (m *Firmware) BatchFetchByResIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []git_querycap_com_tools_datatypes.SFID) ([]Firmware, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ResID").In(values)

	return m.List(db, condition)

}
