// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

// ActivityClassification provides polymorphic access to related types.
// Call the interface's GetActivity() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Activity, *AppendVariableActivity, *AzureDataExplorerCommandActivity, *AzureFunctionActivity, *AzureMLBatchExecutionActivity,
// - *AzureMLExecutePipelineActivity, *AzureMLUpdateResourceActivity, *ControlActivity, *CopyActivity, *CustomActivity, *DataLakeAnalyticsUSQLActivity,
// - *DatabricksNotebookActivity, *DatabricksSparkJarActivity, *DatabricksSparkPythonActivity, *DeleteActivity, *ExecuteDataFlowActivity,
// - *ExecutePipelineActivity, *ExecuteSSISPackageActivity, *ExecuteWranglingDataflowActivity, *ExecutionActivity, *FailActivity,
// - *FilterActivity, *ForEachActivity, *GetMetadataActivity, *HDInsightHiveActivity, *HDInsightMapReduceActivity, *HDInsightPigActivity,
// - *HDInsightSparkActivity, *HDInsightStreamingActivity, *IfConditionActivity, *LookupActivity, *SQLServerStoredProcedureActivity,
// - *ScriptActivity, *SetVariableActivity, *SwitchActivity, *SynapseNotebookActivity, *SynapseSparkJobDefinitionActivity,
// - *UntilActivity, *ValidationActivity, *WaitActivity, *WebActivity, *WebHookActivity
type ActivityClassification interface {
	// GetActivity returns the Activity content of the underlying type.
	GetActivity() *Activity
}

// CompressionReadSettingsClassification provides polymorphic access to related types.
// Call the interface's GetCompressionReadSettings() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CompressionReadSettings, *TarGZipReadSettings, *TarReadSettings, *ZipDeflateReadSettings
type CompressionReadSettingsClassification interface {
	// GetCompressionReadSettings returns the CompressionReadSettings content of the underlying type.
	GetCompressionReadSettings() *CompressionReadSettings
}

// ControlActivityClassification provides polymorphic access to related types.
// Call the interface's GetControlActivity() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AppendVariableActivity, *ControlActivity, *ExecutePipelineActivity, *FailActivity, *FilterActivity, *ForEachActivity,
// - *IfConditionActivity, *SetVariableActivity, *SwitchActivity, *UntilActivity, *ValidationActivity, *WaitActivity, *WebHookActivity
type ControlActivityClassification interface {
	ActivityClassification
	// GetControlActivity returns the ControlActivity content of the underlying type.
	GetControlActivity() *ControlActivity
}

// CopySinkClassification provides polymorphic access to related types.
// Call the interface's GetCopySink() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AvroSink, *AzureBlobFSSink, *AzureDataExplorerSink, *AzureDataLakeStoreSink, *AzureDatabricksDeltaLakeSink, *AzureMySQLSink,
// - *AzurePostgreSQLSink, *AzureQueueSink, *AzureSQLSink, *AzureSearchIndexSink, *AzureTableSink, *BinarySink, *BlobSink,
// - *CommonDataServiceForAppsSink, *CopySink, *CosmosDbMongoDbAPISink, *CosmosDbSQLAPISink, *DelimitedTextSink, *DocumentDbCollectionSink,
// - *DynamicsCrmSink, *DynamicsSink, *FileSystemSink, *IcebergSink, *InformixSink, *JSONSink, *LakeHouseTableSink, *MicrosoftAccessSink,
// - *MongoDbAtlasSink, *MongoDbV2Sink, *OdbcSink, *OracleSink, *OrcSink, *ParquetSink, *RestSink, *SQLDWSink, *SQLMISink,
// - *SQLServerSink, *SQLSink, *SalesforceServiceCloudSink, *SalesforceServiceCloudV2Sink, *SalesforceSink, *SalesforceV2Sink,
// - *SapCloudForCustomerSink, *SnowflakeSink, *SnowflakeV2Sink, *TeradataSink, *WarehouseSink
type CopySinkClassification interface {
	// GetCopySink returns the CopySink content of the underlying type.
	GetCopySink() *CopySink
}

// CopySourceClassification provides polymorphic access to related types.
// Call the interface's GetCopySource() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AmazonMWSSource, *AmazonRdsForOracleSource, *AmazonRdsForSQLServerSource, *AmazonRedshiftSource, *AvroSource, *AzureBlobFSSource,
// - *AzureDataExplorerSource, *AzureDataLakeStoreSource, *AzureDatabricksDeltaLakeSource, *AzureMariaDBSource, *AzureMySQLSource,
// - *AzurePostgreSQLSource, *AzureSQLSource, *AzureTableSource, *BinarySource, *BlobSource, *CassandraSource, *CommonDataServiceForAppsSource,
// - *ConcurSource, *CopySource, *CosmosDbMongoDbAPISource, *CosmosDbSQLAPISource, *CouchbaseSource, *Db2Source, *DelimitedTextSource,
// - *DocumentDbCollectionSource, *DrillSource, *DynamicsAXSource, *DynamicsCrmSource, *DynamicsSource, *EloquaSource, *ExcelSource,
// - *FileSystemSource, *GoogleAdWordsSource, *GoogleBigQuerySource, *GoogleBigQueryV2Source, *GreenplumSource, *HBaseSource,
// - *HTTPSource, *HdfsSource, *HiveSource, *HubspotSource, *ImpalaSource, *InformixSource, *JSONSource, *JiraSource, *LakeHouseTableSource,
// - *MagentoSource, *MariaDBSource, *MarketoSource, *MicrosoftAccessSource, *MongoDbAtlasSource, *MongoDbSource, *MongoDbV2Source,
// - *MySQLSource, *NetezzaSource, *ODataSource, *OdbcSource, *Office365Source, *OracleServiceCloudSource, *OracleSource,
// - *OrcSource, *ParquetSource, *PaypalSource, *PhoenixSource, *PostgreSQLSource, *PostgreSQLV2Source, *PrestoSource, *QuickBooksSource,
// - *RelationalSource, *ResponsysSource, *RestSource, *SQLDWSource, *SQLMISource, *SQLServerSource, *SQLSource, *SalesforceMarketingCloudSource,
// - *SalesforceServiceCloudSource, *SalesforceServiceCloudV2Source, *SalesforceSource, *SalesforceV2Source, *SapBwSource,
// - *SapCloudForCustomerSource, *SapEccSource, *SapHanaSource, *SapOdpSource, *SapOpenHubSource, *SapTableSource, *ServiceNowSource,
// - *ServiceNowV2Source, *SharePointOnlineListSource, *ShopifySource, *SnowflakeSource, *SnowflakeV2Source, *SparkSource,
// - *SquareSource, *SybaseSource, *TabularSource, *TeradataSource, *VerticaSource, *WarehouseSource, *WebSource, *XMLSource,
// - *XeroSource, *ZohoSource
type CopySourceClassification interface {
	// GetCopySource returns the CopySource content of the underlying type.
	GetCopySource() *CopySource
}

// CredentialClassification provides polymorphic access to related types.
// Call the interface's GetCredential() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Credential, *ManagedIdentityCredential, *ServicePrincipalCredential
type CredentialClassification interface {
	// GetCredential returns the Credential content of the underlying type.
	GetCredential() *Credential
}

// CustomSetupBaseClassification provides polymorphic access to related types.
// Call the interface's GetCustomSetupBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzPowerShellSetup, *CmdkeySetup, *ComponentSetup, *CustomSetupBase, *EnvironmentVariableSetup
type CustomSetupBaseClassification interface {
	// GetCustomSetupBase returns the CustomSetupBase content of the underlying type.
	GetCustomSetupBase() *CustomSetupBase
}

// DataFlowClassification provides polymorphic access to related types.
// Call the interface's GetDataFlow() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DataFlow, *Flowlet, *MappingDataFlow, *WranglingDataFlow
type DataFlowClassification interface {
	// GetDataFlow returns the DataFlow content of the underlying type.
	GetDataFlow() *DataFlow
}

// DatasetClassification provides polymorphic access to related types.
// Call the interface's GetDataset() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AmazonMWSObjectDataset, *AmazonRdsForOracleTableDataset, *AmazonRdsForSQLServerTableDataset, *AmazonRedshiftTableDataset,
// - *AmazonS3Dataset, *AvroDataset, *AzureBlobDataset, *AzureBlobFSDataset, *AzureDataExplorerTableDataset, *AzureDataLakeStoreDataset,
// - *AzureDatabricksDeltaLakeDataset, *AzureMariaDBTableDataset, *AzureMySQLTableDataset, *AzurePostgreSQLTableDataset, *AzureSQLDWTableDataset,
// - *AzureSQLMITableDataset, *AzureSQLTableDataset, *AzureSearchIndexDataset, *AzureTableDataset, *BinaryDataset, *CassandraTableDataset,
// - *CommonDataServiceForAppsEntityDataset, *ConcurObjectDataset, *CosmosDbMongoDbAPICollectionDataset, *CosmosDbSQLAPICollectionDataset,
// - *CouchbaseTableDataset, *CustomDataset, *Dataset, *Db2TableDataset, *DelimitedTextDataset, *DocumentDbCollectionDataset,
// - *DrillTableDataset, *DynamicsAXResourceDataset, *DynamicsCrmEntityDataset, *DynamicsEntityDataset, *EloquaObjectDataset,
// - *ExcelDataset, *FileShareDataset, *GoogleAdWordsObjectDataset, *GoogleBigQueryObjectDataset, *GoogleBigQueryV2ObjectDataset,
// - *GreenplumTableDataset, *HBaseObjectDataset, *HTTPDataset, *HiveObjectDataset, *HubspotObjectDataset, *IcebergDataset,
// - *ImpalaObjectDataset, *InformixTableDataset, *JSONDataset, *JiraObjectDataset, *LakeHouseTableDataset, *MagentoObjectDataset,
// - *MariaDBTableDataset, *MarketoObjectDataset, *MicrosoftAccessTableDataset, *MongoDbAtlasCollectionDataset, *MongoDbCollectionDataset,
// - *MongoDbV2CollectionDataset, *MySQLTableDataset, *NetezzaTableDataset, *ODataResourceDataset, *OdbcTableDataset, *Office365Dataset,
// - *OracleServiceCloudObjectDataset, *OracleTableDataset, *OrcDataset, *ParquetDataset, *PaypalObjectDataset, *PhoenixObjectDataset,
// - *PostgreSQLTableDataset, *PostgreSQLV2TableDataset, *PrestoObjectDataset, *QuickBooksObjectDataset, *RelationalTableDataset,
// - *ResponsysObjectDataset, *RestResourceDataset, *SQLServerTableDataset, *SalesforceMarketingCloudObjectDataset, *SalesforceObjectDataset,
// - *SalesforceServiceCloudObjectDataset, *SalesforceServiceCloudV2ObjectDataset, *SalesforceV2ObjectDataset, *SapBwCubeDataset,
// - *SapCloudForCustomerResourceDataset, *SapEccResourceDataset, *SapHanaTableDataset, *SapOdpResourceDataset, *SapOpenHubTableDataset,
// - *SapTableResourceDataset, *ServiceNowObjectDataset, *ServiceNowV2ObjectDataset, *SharePointOnlineListResourceDataset,
// - *ShopifyObjectDataset, *SnowflakeDataset, *SnowflakeV2Dataset, *SparkObjectDataset, *SquareObjectDataset, *SybaseTableDataset,
// - *TeradataTableDataset, *VerticaTableDataset, *WarehouseTableDataset, *WebTableDataset, *XMLDataset, *XeroObjectDataset,
// - *ZohoObjectDataset
type DatasetClassification interface {
	// GetDataset returns the Dataset content of the underlying type.
	GetDataset() *Dataset
}

// DatasetLocationClassification provides polymorphic access to related types.
// Call the interface's GetDatasetLocation() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AmazonS3CompatibleLocation, *AmazonS3Location, *AzureBlobFSLocation, *AzureBlobStorageLocation, *AzureDataLakeStoreLocation,
// - *AzureFileStorageLocation, *DatasetLocation, *FileServerLocation, *FtpServerLocation, *GoogleCloudStorageLocation, *HTTPServerLocation,
// - *HdfsLocation, *LakeHouseLocation, *OracleCloudStorageLocation, *SftpLocation
type DatasetLocationClassification interface {
	// GetDatasetLocation returns the DatasetLocation content of the underlying type.
	GetDatasetLocation() *DatasetLocation
}

// DatasetStorageFormatClassification provides polymorphic access to related types.
// Call the interface's GetDatasetStorageFormat() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AvroFormat, *DatasetStorageFormat, *JSONFormat, *OrcFormat, *ParquetFormat, *TextFormat
type DatasetStorageFormatClassification interface {
	// GetDatasetStorageFormat returns the DatasetStorageFormat content of the underlying type.
	GetDatasetStorageFormat() *DatasetStorageFormat
}

// DependencyReferenceClassification provides polymorphic access to related types.
// Call the interface's GetDependencyReference() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DependencyReference, *SelfDependencyTumblingWindowTriggerReference, *TriggerDependencyReference, *TumblingWindowTriggerDependencyReference
type DependencyReferenceClassification interface {
	// GetDependencyReference returns the DependencyReference content of the underlying type.
	GetDependencyReference() *DependencyReference
}

// ExecutionActivityClassification provides polymorphic access to related types.
// Call the interface's GetExecutionActivity() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureDataExplorerCommandActivity, *AzureFunctionActivity, *AzureMLBatchExecutionActivity, *AzureMLExecutePipelineActivity,
// - *AzureMLUpdateResourceActivity, *CopyActivity, *CustomActivity, *DataLakeAnalyticsUSQLActivity, *DatabricksNotebookActivity,
// - *DatabricksSparkJarActivity, *DatabricksSparkPythonActivity, *DeleteActivity, *ExecuteDataFlowActivity, *ExecuteSSISPackageActivity,
// - *ExecutionActivity, *GetMetadataActivity, *HDInsightHiveActivity, *HDInsightMapReduceActivity, *HDInsightPigActivity,
// - *HDInsightSparkActivity, *HDInsightStreamingActivity, *LookupActivity, *SQLServerStoredProcedureActivity, *ScriptActivity,
// - *SynapseNotebookActivity, *SynapseSparkJobDefinitionActivity, *WebActivity
type ExecutionActivityClassification interface {
	ActivityClassification
	// GetExecutionActivity returns the ExecutionActivity content of the underlying type.
	GetExecutionActivity() *ExecutionActivity
}

// ExportSettingsClassification provides polymorphic access to related types.
// Call the interface's GetExportSettings() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureDatabricksDeltaLakeExportCommand, *ExportSettings, *SnowflakeExportCopyCommand
type ExportSettingsClassification interface {
	// GetExportSettings returns the ExportSettings content of the underlying type.
	GetExportSettings() *ExportSettings
}

// FactoryRepoConfigurationClassification provides polymorphic access to related types.
// Call the interface's GetFactoryRepoConfiguration() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *FactoryGitHubConfiguration, *FactoryRepoConfiguration, *FactoryVSTSConfiguration
type FactoryRepoConfigurationClassification interface {
	// GetFactoryRepoConfiguration returns the FactoryRepoConfiguration content of the underlying type.
	GetFactoryRepoConfiguration() *FactoryRepoConfiguration
}

// FormatReadSettingsClassification provides polymorphic access to related types.
// Call the interface's GetFormatReadSettings() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BinaryReadSettings, *DelimitedTextReadSettings, *FormatReadSettings, *JSONReadSettings, *ParquetReadSettings, *XMLReadSettings
type FormatReadSettingsClassification interface {
	// GetFormatReadSettings returns the FormatReadSettings content of the underlying type.
	GetFormatReadSettings() *FormatReadSettings
}

// FormatWriteSettingsClassification provides polymorphic access to related types.
// Call the interface's GetFormatWriteSettings() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AvroWriteSettings, *DelimitedTextWriteSettings, *FormatWriteSettings, *IcebergWriteSettings, *JSONWriteSettings, *OrcWriteSettings,
// - *ParquetWriteSettings
type FormatWriteSettingsClassification interface {
	// GetFormatWriteSettings returns the FormatWriteSettings content of the underlying type.
	GetFormatWriteSettings() *FormatWriteSettings
}

// ImportSettingsClassification provides polymorphic access to related types.
// Call the interface's GetImportSettings() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureDatabricksDeltaLakeImportCommand, *ImportSettings, *SnowflakeImportCopyCommand, *TeradataImportCommand
type ImportSettingsClassification interface {
	// GetImportSettings returns the ImportSettings content of the underlying type.
	GetImportSettings() *ImportSettings
}

// IntegrationRuntimeClassification provides polymorphic access to related types.
// Call the interface's GetIntegrationRuntime() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *IntegrationRuntime, *ManagedIntegrationRuntime, *SelfHostedIntegrationRuntime
type IntegrationRuntimeClassification interface {
	// GetIntegrationRuntime returns the IntegrationRuntime content of the underlying type.
	GetIntegrationRuntime() *IntegrationRuntime
}

// IntegrationRuntimeStatusClassification provides polymorphic access to related types.
// Call the interface's GetIntegrationRuntimeStatus() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *IntegrationRuntimeStatus, *ManagedIntegrationRuntimeStatus, *SelfHostedIntegrationRuntimeStatus
type IntegrationRuntimeStatusClassification interface {
	// GetIntegrationRuntimeStatus returns the IntegrationRuntimeStatus content of the underlying type.
	GetIntegrationRuntimeStatus() *IntegrationRuntimeStatus
}

// LinkedIntegrationRuntimeTypeClassification provides polymorphic access to related types.
// Call the interface's GetLinkedIntegrationRuntimeType() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *LinkedIntegrationRuntimeKeyAuthorization, *LinkedIntegrationRuntimeRbacAuthorization, *LinkedIntegrationRuntimeType
type LinkedIntegrationRuntimeTypeClassification interface {
	// GetLinkedIntegrationRuntimeType returns the LinkedIntegrationRuntimeType content of the underlying type.
	GetLinkedIntegrationRuntimeType() *LinkedIntegrationRuntimeType
}

// LinkedServiceClassification provides polymorphic access to related types.
// Call the interface's GetLinkedService() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AmazonMWSLinkedService, *AmazonRdsForOracleLinkedService, *AmazonRdsForSQLServerLinkedService, *AmazonRedshiftLinkedService,
// - *AmazonS3CompatibleLinkedService, *AmazonS3LinkedService, *AppFiguresLinkedService, *AsanaLinkedService, *AzureBatchLinkedService,
// - *AzureBlobFSLinkedService, *AzureBlobStorageLinkedService, *AzureDataExplorerLinkedService, *AzureDataLakeAnalyticsLinkedService,
// - *AzureDataLakeStoreLinkedService, *AzureDatabricksDeltaLakeLinkedService, *AzureDatabricksLinkedService, *AzureFileStorageLinkedService,
// - *AzureFunctionLinkedService, *AzureKeyVaultLinkedService, *AzureMLLinkedService, *AzureMLServiceLinkedService, *AzureMariaDBLinkedService,
// - *AzureMySQLLinkedService, *AzurePostgreSQLLinkedService, *AzureSQLDWLinkedService, *AzureSQLDatabaseLinkedService, *AzureSQLMILinkedService,
// - *AzureSearchLinkedService, *AzureStorageLinkedService, *AzureSynapseArtifactsLinkedService, *AzureTableStorageLinkedService,
// - *CassandraLinkedService, *CommonDataServiceForAppsLinkedService, *ConcurLinkedService, *CosmosDbLinkedService, *CosmosDbMongoDbAPILinkedService,
// - *CouchbaseLinkedService, *CustomDataSourceLinkedService, *DataworldLinkedService, *Db2LinkedService, *DrillLinkedService,
// - *DynamicsAXLinkedService, *DynamicsCrmLinkedService, *DynamicsLinkedService, *EloquaLinkedService, *FileServerLinkedService,
// - *FtpServerLinkedService, *GoogleAdWordsLinkedService, *GoogleBigQueryLinkedService, *GoogleBigQueryV2LinkedService, *GoogleCloudStorageLinkedService,
// - *GoogleSheetsLinkedService, *GreenplumLinkedService, *HBaseLinkedService, *HDInsightLinkedService, *HDInsightOnDemandLinkedService,
// - *HTTPLinkedService, *HdfsLinkedService, *HiveLinkedService, *HubspotLinkedService, *ImpalaLinkedService, *InformixLinkedService,
// - *JiraLinkedService, *LakeHouseLinkedService, *LinkedService, *MagentoLinkedService, *MariaDBLinkedService, *MarketoLinkedService,
// - *MicrosoftAccessLinkedService, *MongoDbAtlasLinkedService, *MongoDbLinkedService, *MongoDbV2LinkedService, *MySQLLinkedService,
// - *NetezzaLinkedService, *ODataLinkedService, *OdbcLinkedService, *Office365LinkedService, *OracleCloudStorageLinkedService,
// - *OracleLinkedService, *OracleServiceCloudLinkedService, *PaypalLinkedService, *PhoenixLinkedService, *PostgreSQLLinkedService,
// - *PostgreSQLV2LinkedService, *PrestoLinkedService, *QuickBooksLinkedService, *QuickbaseLinkedService, *ResponsysLinkedService,
// - *RestServiceLinkedService, *SQLServerLinkedService, *SalesforceLinkedService, *SalesforceMarketingCloudLinkedService,
// - *SalesforceServiceCloudLinkedService, *SalesforceServiceCloudV2LinkedService, *SalesforceV2LinkedService, *SapBWLinkedService,
// - *SapCloudForCustomerLinkedService, *SapEccLinkedService, *SapHanaLinkedService, *SapOdpLinkedService, *SapOpenHubLinkedService,
// - *SapTableLinkedService, *ServiceNowLinkedService, *ServiceNowV2LinkedService, *SftpServerLinkedService, *SharePointOnlineListLinkedService,
// - *ShopifyLinkedService, *SmartsheetLinkedService, *SnowflakeLinkedService, *SnowflakeV2LinkedService, *SparkLinkedService,
// - *SquareLinkedService, *SybaseLinkedService, *TeamDeskLinkedService, *TeradataLinkedService, *TwilioLinkedService, *VerticaLinkedService,
// - *WarehouseLinkedService, *WebLinkedService, *XeroLinkedService, *ZendeskLinkedService, *ZohoLinkedService
type LinkedServiceClassification interface {
	// GetLinkedService returns the LinkedService content of the underlying type.
	GetLinkedService() *LinkedService
}

// MultiplePipelineTriggerClassification provides polymorphic access to related types.
// Call the interface's GetMultiplePipelineTrigger() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BlobEventsTrigger, *BlobTrigger, *CustomEventsTrigger, *MultiplePipelineTrigger, *ScheduleTrigger
type MultiplePipelineTriggerClassification interface {
	TriggerClassification
	// GetMultiplePipelineTrigger returns the MultiplePipelineTrigger content of the underlying type.
	GetMultiplePipelineTrigger() *MultiplePipelineTrigger
}

// SecretBaseClassification provides polymorphic access to related types.
// Call the interface's GetSecretBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureKeyVaultSecretReference, *SecretBase, *SecureString
type SecretBaseClassification interface {
	// GetSecretBase returns the SecretBase content of the underlying type.
	GetSecretBase() *SecretBase
}

// SsisObjectMetadataClassification provides polymorphic access to related types.
// Call the interface's GetSsisObjectMetadata() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *SsisEnvironment, *SsisFolder, *SsisObjectMetadata, *SsisPackage, *SsisProject
type SsisObjectMetadataClassification interface {
	// GetSsisObjectMetadata returns the SsisObjectMetadata content of the underlying type.
	GetSsisObjectMetadata() *SsisObjectMetadata
}

// StoreReadSettingsClassification provides polymorphic access to related types.
// Call the interface's GetStoreReadSettings() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AmazonS3CompatibleReadSettings, *AmazonS3ReadSettings, *AzureBlobFSReadSettings, *AzureBlobStorageReadSettings, *AzureDataLakeStoreReadSettings,
// - *AzureFileStorageReadSettings, *FileServerReadSettings, *FtpReadSettings, *GoogleCloudStorageReadSettings, *HTTPReadSettings,
// - *HdfsReadSettings, *LakeHouseReadSettings, *OracleCloudStorageReadSettings, *SftpReadSettings, *StoreReadSettings
type StoreReadSettingsClassification interface {
	// GetStoreReadSettings returns the StoreReadSettings content of the underlying type.
	GetStoreReadSettings() *StoreReadSettings
}

// StoreWriteSettingsClassification provides polymorphic access to related types.
// Call the interface's GetStoreWriteSettings() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureBlobFSWriteSettings, *AzureBlobStorageWriteSettings, *AzureDataLakeStoreWriteSettings, *AzureFileStorageWriteSettings,
// - *FileServerWriteSettings, *LakeHouseWriteSettings, *SftpWriteSettings, *StoreWriteSettings
type StoreWriteSettingsClassification interface {
	// GetStoreWriteSettings returns the StoreWriteSettings content of the underlying type.
	GetStoreWriteSettings() *StoreWriteSettings
}

// TabularSourceClassification provides polymorphic access to related types.
// Call the interface's GetTabularSource() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AmazonMWSSource, *AmazonRdsForSQLServerSource, *AmazonRedshiftSource, *AzureMariaDBSource, *AzureMySQLSource, *AzurePostgreSQLSource,
// - *AzureSQLSource, *AzureTableSource, *CassandraSource, *ConcurSource, *CouchbaseSource, *Db2Source, *DrillSource, *DynamicsAXSource,
// - *EloquaSource, *GoogleAdWordsSource, *GoogleBigQuerySource, *GoogleBigQueryV2Source, *GreenplumSource, *HBaseSource,
// - *HiveSource, *HubspotSource, *ImpalaSource, *InformixSource, *JiraSource, *MagentoSource, *MariaDBSource, *MarketoSource,
// - *MySQLSource, *NetezzaSource, *OdbcSource, *OracleServiceCloudSource, *PaypalSource, *PhoenixSource, *PostgreSQLSource,
// - *PostgreSQLV2Source, *PrestoSource, *QuickBooksSource, *ResponsysSource, *SQLDWSource, *SQLMISource, *SQLServerSource,
// - *SQLSource, *SalesforceMarketingCloudSource, *SalesforceSource, *SalesforceV2Source, *SapBwSource, *SapCloudForCustomerSource,
// - *SapEccSource, *SapHanaSource, *SapOdpSource, *SapOpenHubSource, *SapTableSource, *ServiceNowSource, *ServiceNowV2Source,
// - *ShopifySource, *SparkSource, *SquareSource, *SybaseSource, *TabularSource, *TeradataSource, *VerticaSource, *WarehouseSource,
// - *XeroSource, *ZohoSource
type TabularSourceClassification interface {
	CopySourceClassification
	// GetTabularSource returns the TabularSource content of the underlying type.
	GetTabularSource() *TabularSource
}

// TriggerClassification provides polymorphic access to related types.
// Call the interface's GetTrigger() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BlobEventsTrigger, *BlobTrigger, *ChainingTrigger, *CustomEventsTrigger, *MultiplePipelineTrigger, *RerunTumblingWindowTrigger,
// - *ScheduleTrigger, *Trigger, *TumblingWindowTrigger
type TriggerClassification interface {
	// GetTrigger returns the Trigger content of the underlying type.
	GetTrigger() *Trigger
}

// TriggerDependencyReferenceClassification provides polymorphic access to related types.
// Call the interface's GetTriggerDependencyReference() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *TriggerDependencyReference, *TumblingWindowTriggerDependencyReference
type TriggerDependencyReferenceClassification interface {
	DependencyReferenceClassification
	// GetTriggerDependencyReference returns the TriggerDependencyReference content of the underlying type.
	GetTriggerDependencyReference() *TriggerDependencyReference
}

// WebLinkedServiceTypePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetWebLinkedServiceTypeProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *WebAnonymousAuthentication, *WebBasicAuthentication, *WebClientCertificateAuthentication, *WebLinkedServiceTypeProperties
type WebLinkedServiceTypePropertiesClassification interface {
	// GetWebLinkedServiceTypeProperties returns the WebLinkedServiceTypeProperties content of the underlying type.
	GetWebLinkedServiceTypeProperties() *WebLinkedServiceTypeProperties
}
