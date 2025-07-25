// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

import "encoding/json"

func unmarshalCommandPropertiesClassification(rawMsg json.RawMessage) (CommandPropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CommandPropertiesClassification
	switch m["commandType"] {
	case string(CommandTypeMigrateSQLServerAzureDbSQLMiComplete):
		b = &MigrateMISyncCompleteCommandProperties{}
	case string(CommandTypeMigrateSyncCompleteDatabase):
		b = &MigrateSyncCompleteCommandProperties{}
	case string(CommandTypeCancel):
		b = &MongoDbCancelCommand{}
	case string(CommandTypeFinish):
		b = &MongoDbFinishCommand{}
	case string(CommandTypeRestart):
		b = &MongoDbRestartCommand{}
	default:
		b = &CommandProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalCommandPropertiesClassificationArray(rawMsg json.RawMessage) ([]CommandPropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]CommandPropertiesClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalCommandPropertiesClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalConnectToSourceSQLServerTaskOutputClassification(rawMsg json.RawMessage) (ConnectToSourceSQLServerTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ConnectToSourceSQLServerTaskOutputClassification
	switch m["resultType"] {
	case "AgentJobLevelOutput":
		b = &ConnectToSourceSQLServerTaskOutputAgentJobLevel{}
	case "DatabaseLevelOutput":
		b = &ConnectToSourceSQLServerTaskOutputDatabaseLevel{}
	case "LoginLevelOutput":
		b = &ConnectToSourceSQLServerTaskOutputLoginLevel{}
	case "TaskLevelOutput":
		b = &ConnectToSourceSQLServerTaskOutputTaskLevel{}
	default:
		b = &ConnectToSourceSQLServerTaskOutput{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalConnectToSourceSQLServerTaskOutputClassificationArray(rawMsg json.RawMessage) ([]ConnectToSourceSQLServerTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ConnectToSourceSQLServerTaskOutputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalConnectToSourceSQLServerTaskOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalConnectionInfoClassification(rawMsg json.RawMessage) (ConnectionInfoClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ConnectionInfoClassification
	switch m["type"] {
	case "MiSqlConnectionInfo":
		b = &MiSQLConnectionInfo{}
	case "MongoDbConnectionInfo":
		b = &MongoDbConnectionInfo{}
	case "MySqlConnectionInfo":
		b = &MySQLConnectionInfo{}
	case "OracleConnectionInfo":
		b = &OracleConnectionInfo{}
	case "PostgreSqlConnectionInfo":
		b = &PostgreSQLConnectionInfo{}
	case "SqlConnectionInfo":
		b = &SQLConnectionInfo{}
	default:
		b = &ConnectionInfo{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDatabaseMigrationBasePropertiesClassification(rawMsg json.RawMessage) (DatabaseMigrationBasePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DatabaseMigrationBasePropertiesClassification
	switch m["kind"] {
	case "DatabaseMigrationProperties":
		b = &DatabaseMigrationProperties{}
	case string(ResourceTypeMongoToCosmosDbMongo):
		b = &DatabaseMigrationPropertiesCosmosDbMongo{}
	case string(ResourceTypeSQLDb):
		b = &DatabaseMigrationPropertiesSQLDb{}
	case string(ResourceTypeSQLMi):
		b = &DatabaseMigrationPropertiesSQLMi{}
	case string(ResourceTypeSQLVM):
		b = &DatabaseMigrationPropertiesSQLVM{}
	default:
		b = &DatabaseMigrationBaseProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDatabaseMigrationPropertiesClassification(rawMsg json.RawMessage) (DatabaseMigrationPropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DatabaseMigrationPropertiesClassification
	switch m["kind"] {
	case string(ResourceTypeSQLDb):
		b = &DatabaseMigrationPropertiesSQLDb{}
	case string(ResourceTypeSQLMi):
		b = &DatabaseMigrationPropertiesSQLMi{}
	case string(ResourceTypeSQLVM):
		b = &DatabaseMigrationPropertiesSQLVM{}
	default:
		b = &DatabaseMigrationProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMigrateMySQLAzureDbForMySQLOfflineTaskOutputClassification(rawMsg json.RawMessage) (MigrateMySQLAzureDbForMySQLOfflineTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MigrateMySQLAzureDbForMySQLOfflineTaskOutputClassification
	switch m["resultType"] {
	case "DatabaseLevelOutput":
		b = &MigrateMySQLAzureDbForMySQLOfflineTaskOutputDatabaseLevel{}
	case "ErrorOutput":
		b = &MigrateMySQLAzureDbForMySQLOfflineTaskOutputError{}
	case "MigrationLevelOutput":
		b = &MigrateMySQLAzureDbForMySQLOfflineTaskOutputMigrationLevel{}
	case "TableLevelOutput":
		b = &MigrateMySQLAzureDbForMySQLOfflineTaskOutputTableLevel{}
	default:
		b = &MigrateMySQLAzureDbForMySQLOfflineTaskOutput{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMigrateMySQLAzureDbForMySQLOfflineTaskOutputClassificationArray(rawMsg json.RawMessage) ([]MigrateMySQLAzureDbForMySQLOfflineTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MigrateMySQLAzureDbForMySQLOfflineTaskOutputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMigrateMySQLAzureDbForMySQLOfflineTaskOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalMigrateMySQLAzureDbForMySQLSyncTaskOutputClassification(rawMsg json.RawMessage) (MigrateMySQLAzureDbForMySQLSyncTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MigrateMySQLAzureDbForMySQLSyncTaskOutputClassification
	switch m["resultType"] {
	case "DatabaseLevelErrorOutput":
		b = &MigrateMySQLAzureDbForMySQLSyncTaskOutputDatabaseError{}
	case "DatabaseLevelOutput":
		b = &MigrateMySQLAzureDbForMySQLSyncTaskOutputDatabaseLevel{}
	case "ErrorOutput":
		b = &MigrateMySQLAzureDbForMySQLSyncTaskOutputError{}
	case "MigrationLevelOutput":
		b = &MigrateMySQLAzureDbForMySQLSyncTaskOutputMigrationLevel{}
	case "TableLevelOutput":
		b = &MigrateMySQLAzureDbForMySQLSyncTaskOutputTableLevel{}
	default:
		b = &MigrateMySQLAzureDbForMySQLSyncTaskOutput{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMigrateMySQLAzureDbForMySQLSyncTaskOutputClassificationArray(rawMsg json.RawMessage) ([]MigrateMySQLAzureDbForMySQLSyncTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MigrateMySQLAzureDbForMySQLSyncTaskOutputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMigrateMySQLAzureDbForMySQLSyncTaskOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalMigrateOracleAzureDbPostgreSQLSyncTaskOutputClassification(rawMsg json.RawMessage) (MigrateOracleAzureDbPostgreSQLSyncTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MigrateOracleAzureDbPostgreSQLSyncTaskOutputClassification
	switch m["resultType"] {
	case "DatabaseLevelErrorOutput":
		b = &MigrateOracleAzureDbPostgreSQLSyncTaskOutputDatabaseError{}
	case "DatabaseLevelOutput":
		b = &MigrateOracleAzureDbPostgreSQLSyncTaskOutputDatabaseLevel{}
	case "ErrorOutput":
		b = &MigrateOracleAzureDbPostgreSQLSyncTaskOutputError{}
	case "MigrationLevelOutput":
		b = &MigrateOracleAzureDbPostgreSQLSyncTaskOutputMigrationLevel{}
	case "TableLevelOutput":
		b = &MigrateOracleAzureDbPostgreSQLSyncTaskOutputTableLevel{}
	default:
		b = &MigrateOracleAzureDbPostgreSQLSyncTaskOutput{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMigrateOracleAzureDbPostgreSQLSyncTaskOutputClassificationArray(rawMsg json.RawMessage) ([]MigrateOracleAzureDbPostgreSQLSyncTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MigrateOracleAzureDbPostgreSQLSyncTaskOutputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMigrateOracleAzureDbPostgreSQLSyncTaskOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalMigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputClassification(rawMsg json.RawMessage) (MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputClassification
	switch m["resultType"] {
	case "DatabaseLevelErrorOutput":
		b = &MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputDatabaseError{}
	case "DatabaseLevelOutput":
		b = &MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputDatabaseLevel{}
	case "ErrorOutput":
		b = &MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputError{}
	case "MigrationLevelOutput":
		b = &MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputMigrationLevel{}
	case "TableLevelOutput":
		b = &MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputTableLevel{}
	default:
		b = &MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutput{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputClassificationArray(rawMsg json.RawMessage) ([]MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalMigrateSQLServerSQLDbSyncTaskOutputClassification(rawMsg json.RawMessage) (MigrateSQLServerSQLDbSyncTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MigrateSQLServerSQLDbSyncTaskOutputClassification
	switch m["resultType"] {
	case "DatabaseLevelErrorOutput":
		b = &MigrateSQLServerSQLDbSyncTaskOutputDatabaseError{}
	case "DatabaseLevelOutput":
		b = &MigrateSQLServerSQLDbSyncTaskOutputDatabaseLevel{}
	case "ErrorOutput":
		b = &MigrateSQLServerSQLDbSyncTaskOutputError{}
	case "MigrationLevelOutput":
		b = &MigrateSQLServerSQLDbSyncTaskOutputMigrationLevel{}
	case "TableLevelOutput":
		b = &MigrateSQLServerSQLDbSyncTaskOutputTableLevel{}
	default:
		b = &MigrateSQLServerSQLDbSyncTaskOutput{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMigrateSQLServerSQLDbSyncTaskOutputClassificationArray(rawMsg json.RawMessage) ([]MigrateSQLServerSQLDbSyncTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MigrateSQLServerSQLDbSyncTaskOutputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMigrateSQLServerSQLDbSyncTaskOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalMigrateSQLServerSQLDbTaskOutputClassification(rawMsg json.RawMessage) (MigrateSQLServerSQLDbTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MigrateSQLServerSQLDbTaskOutputClassification
	switch m["resultType"] {
	case "DatabaseLevelOutput":
		b = &MigrateSQLServerSQLDbTaskOutputDatabaseLevel{}
	case "ErrorOutput":
		b = &MigrateSQLServerSQLDbTaskOutputError{}
	case "MigrationDatabaseLevelValidationOutput":
		b = &MigrateSQLServerSQLDbTaskOutputDatabaseLevelValidationResult{}
	case "MigrationLevelOutput":
		b = &MigrateSQLServerSQLDbTaskOutputMigrationLevel{}
	case "MigrationValidationOutput":
		b = &MigrateSQLServerSQLDbTaskOutputValidationResult{}
	case "TableLevelOutput":
		b = &MigrateSQLServerSQLDbTaskOutputTableLevel{}
	default:
		b = &MigrateSQLServerSQLDbTaskOutput{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMigrateSQLServerSQLDbTaskOutputClassificationArray(rawMsg json.RawMessage) ([]MigrateSQLServerSQLDbTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MigrateSQLServerSQLDbTaskOutputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMigrateSQLServerSQLDbTaskOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalMigrateSQLServerSQLMISyncTaskOutputClassification(rawMsg json.RawMessage) (MigrateSQLServerSQLMISyncTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MigrateSQLServerSQLMISyncTaskOutputClassification
	switch m["resultType"] {
	case "DatabaseLevelOutput":
		b = &MigrateSQLServerSQLMISyncTaskOutputDatabaseLevel{}
	case "ErrorOutput":
		b = &MigrateSQLServerSQLMISyncTaskOutputError{}
	case "MigrationLevelOutput":
		b = &MigrateSQLServerSQLMISyncTaskOutputMigrationLevel{}
	default:
		b = &MigrateSQLServerSQLMISyncTaskOutput{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMigrateSQLServerSQLMISyncTaskOutputClassificationArray(rawMsg json.RawMessage) ([]MigrateSQLServerSQLMISyncTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MigrateSQLServerSQLMISyncTaskOutputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMigrateSQLServerSQLMISyncTaskOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalMigrateSQLServerSQLMITaskOutputClassification(rawMsg json.RawMessage) (MigrateSQLServerSQLMITaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MigrateSQLServerSQLMITaskOutputClassification
	switch m["resultType"] {
	case "AgentJobLevelOutput":
		b = &MigrateSQLServerSQLMITaskOutputAgentJobLevel{}
	case "DatabaseLevelOutput":
		b = &MigrateSQLServerSQLMITaskOutputDatabaseLevel{}
	case "ErrorOutput":
		b = &MigrateSQLServerSQLMITaskOutputError{}
	case "LoginLevelOutput":
		b = &MigrateSQLServerSQLMITaskOutputLoginLevel{}
	case "MigrationLevelOutput":
		b = &MigrateSQLServerSQLMITaskOutputMigrationLevel{}
	default:
		b = &MigrateSQLServerSQLMITaskOutput{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMigrateSQLServerSQLMITaskOutputClassificationArray(rawMsg json.RawMessage) ([]MigrateSQLServerSQLMITaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MigrateSQLServerSQLMITaskOutputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMigrateSQLServerSQLMITaskOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalMigrateSchemaSQLServerSQLDbTaskOutputClassification(rawMsg json.RawMessage) (MigrateSchemaSQLServerSQLDbTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MigrateSchemaSQLServerSQLDbTaskOutputClassification
	switch m["resultType"] {
	case "DatabaseLevelOutput":
		b = &MigrateSchemaSQLServerSQLDbTaskOutputDatabaseLevel{}
	case "ErrorOutput":
		b = &MigrateSchemaSQLTaskOutputError{}
	case "MigrationLevelOutput":
		b = &MigrateSchemaSQLServerSQLDbTaskOutputMigrationLevel{}
	case "SchemaErrorOutput":
		b = &MigrateSchemaSQLServerSQLDbTaskOutputError{}
	default:
		b = &MigrateSchemaSQLServerSQLDbTaskOutput{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMigrateSchemaSQLServerSQLDbTaskOutputClassificationArray(rawMsg json.RawMessage) ([]MigrateSchemaSQLServerSQLDbTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MigrateSchemaSQLServerSQLDbTaskOutputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMigrateSchemaSQLServerSQLDbTaskOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalMigrateSsisTaskOutputClassification(rawMsg json.RawMessage) (MigrateSsisTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MigrateSsisTaskOutputClassification
	switch m["resultType"] {
	case "MigrationLevelOutput":
		b = &MigrateSsisTaskOutputMigrationLevel{}
	case "SsisProjectLevelOutput":
		b = &MigrateSsisTaskOutputProjectLevel{}
	default:
		b = &MigrateSsisTaskOutput{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMigrateSsisTaskOutputClassificationArray(rawMsg json.RawMessage) ([]MigrateSsisTaskOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MigrateSsisTaskOutputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMigrateSsisTaskOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalMongoDbProgressClassification(rawMsg json.RawMessage) (MongoDbProgressClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MongoDbProgressClassification
	switch m["resultType"] {
	case string(MongoDbProgressResultTypeCollection):
		b = &MongoDbCollectionProgress{}
	case string(MongoDbProgressResultTypeDatabase):
		b = &MongoDbDatabaseProgress{}
	case string(MongoDbProgressResultTypeMigration):
		b = &MongoDbMigrationProgress{}
	default:
		b = &MongoDbProgress{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMongoDbProgressClassificationArray(rawMsg json.RawMessage) ([]MongoDbProgressClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MongoDbProgressClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMongoDbProgressClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalProjectTaskPropertiesClassification(rawMsg json.RawMessage) (ProjectTaskPropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProjectTaskPropertiesClassification
	switch m["taskType"] {
	case string(TaskTypeConnectMongoDb):
		b = &ConnectToMongoDbTaskProperties{}
	case string(TaskTypeConnectToSourceMySQL):
		b = &ConnectToSourceMySQLTaskProperties{}
	case string(TaskTypeConnectToSourceOracleSync):
		b = &ConnectToSourceOracleSyncTaskProperties{}
	case string(TaskTypeConnectToSourcePostgreSQLSync):
		b = &ConnectToSourcePostgreSQLSyncTaskProperties{}
	case string(TaskTypeConnectToSourceSQLServer):
		b = &ConnectToSourceSQLServerTaskProperties{}
	case string(TaskTypeConnectToSourceSQLServerSync):
		b = &ConnectToSourceSQLServerSyncTaskProperties{}
	case string(TaskTypeConnectToTargetAzureDbForMySQL):
		b = &ConnectToTargetAzureDbForMySQLTaskProperties{}
	case string(TaskTypeConnectToTargetAzureDbForPostgreSQLSync):
		b = &ConnectToTargetAzureDbForPostgreSQLSyncTaskProperties{}
	case string(TaskTypeConnectToTargetAzureSQLDbMI):
		b = &ConnectToTargetSQLMITaskProperties{}
	case string(TaskTypeConnectToTargetAzureSQLDbMISyncLRS):
		b = &ConnectToTargetSQLMISyncTaskProperties{}
	case string(TaskTypeConnectToTargetOracleAzureDbForPostgreSQLSync):
		b = &ConnectToTargetOracleAzureDbForPostgreSQLSyncTaskProperties{}
	case string(TaskTypeConnectToTargetSQLDb):
		b = &ConnectToTargetSQLDbTaskProperties{}
	case string(TaskTypeConnectToTargetSQLDbSync):
		b = &ConnectToTargetSQLDbSyncTaskProperties{}
	case string(TaskTypeGetTDECertificatesSQL):
		b = &GetTdeCertificatesSQLTaskProperties{}
	case string(TaskTypeGetUserTablesAzureSQLDbSync):
		b = &GetUserTablesSQLSyncTaskProperties{}
	case string(TaskTypeGetUserTablesSQL):
		b = &GetUserTablesSQLTaskProperties{}
	case string(TaskTypeGetUserTablesMySQL):
		b = &GetUserTablesMySQLTaskProperties{}
	case string(TaskTypeGetUserTablesOracle):
		b = &GetUserTablesOracleTaskProperties{}
	case string(TaskTypeGetUserTablesPostgreSQL):
		b = &GetUserTablesPostgreSQLTaskProperties{}
	case string(TaskTypeMigrateMongoDb):
		b = &MigrateMongoDbTaskProperties{}
	case string(TaskTypeMigrateMySQLAzureDbForMySQL):
		b = &MigrateMySQLAzureDbForMySQLOfflineTaskProperties{}
	case string(TaskTypeMigrateMySQLAzureDbForMySQLSync):
		b = &MigrateMySQLAzureDbForMySQLSyncTaskProperties{}
	case string(TaskTypeMigrateOracleAzureDbForPostgreSQLSync):
		b = &MigrateOracleAzureDbForPostgreSQLSyncTaskProperties{}
	case string(TaskTypeMigratePostgreSQLAzureDbForPostgreSQLSyncV2):
		b = &MigratePostgreSQLAzureDbForPostgreSQLSyncTaskProperties{}
	case string(TaskTypeMigrateSQLServerAzureSQLDbSync):
		b = &MigrateSQLServerSQLDbSyncTaskProperties{}
	case string(TaskTypeMigrateSQLServerAzureSQLDbMI):
		b = &MigrateSQLServerSQLMITaskProperties{}
	case string(TaskTypeMigrateSQLServerAzureSQLDbMISyncLRS):
		b = &MigrateSQLServerSQLMISyncTaskProperties{}
	case string(TaskTypeMigrateSQLServerSQLDb):
		b = &MigrateSQLServerSQLDbTaskProperties{}
	case string(TaskTypeMigrateSsis):
		b = &MigrateSsisTaskProperties{}
	case string(TaskTypeMigrateSchemaSQLServerSQLDb):
		b = &MigrateSchemaSQLServerSQLDbTaskProperties{}
	case string(TaskTypeServiceCheckOCI):
		b = &CheckOCIDriverTaskProperties{}
	case string(TaskTypeServiceInstallOCI):
		b = &InstallOCIDriverTaskProperties{}
	case string(TaskTypeServiceUploadOCI):
		b = &UploadOCIDriverTaskProperties{}
	case string(TaskTypeValidateMongoDb):
		b = &ValidateMongoDbTaskProperties{}
	case string(TaskTypeValidateOracleAzureDbPostgreSQLSync):
		b = &ValidateOracleAzureDbForPostgreSQLSyncTaskProperties{}
	case string(TaskTypeValidateMigrationInputSQLServerAzureSQLDbMI):
		b = &ValidateMigrationInputSQLServerSQLMITaskProperties{}
	case string(TaskTypeValidateMigrationInputSQLServerAzureSQLDbMISyncLRS):
		b = &ValidateMigrationInputSQLServerSQLMISyncTaskProperties{}
	case string(TaskTypeValidateMigrationInputSQLServerSQLDbSync):
		b = &ValidateMigrationInputSQLServerSQLDbSyncTaskProperties{}
	default:
		b = &ProjectTaskProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
