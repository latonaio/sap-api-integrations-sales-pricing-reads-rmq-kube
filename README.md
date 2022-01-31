# sap-api-integrations-sales-pricing-reads-rmq-kube
sap-api-integrations-sales-pricing-reads-rmq-kube は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で販売価格を取得するマイクロサービスです。    
sap-api-integrations-sales-pricing-reads-rmq-kube には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-sales-pricing-reads-rmq-kube は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_SLSPRCGCONDITIONRECORD_SRV_0001/overview   

## 動作環境  
sap-api-integrations-sales-pricing-reads-rmq-kube は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  
・ RabbitMQ on Kubernetes  
・ RabbitMQ Client   

## クラウド環境での利用
sap-api-integrations-sales-pricing-reads-rmq-kube は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## RabbitMQ からの JSON Input

sap-api-integrations-sales-pricing-reads-rmq-kube は、Inputとして、RabbitMQ からのメッセージをJSON形式で受け取ります。 
Input の サンプルJSON は、Inputs フォルダ内にあります。  

## RabbitMQ からのメッセージ受信による イベントドリヴン の ランタイム実行

sap-api-integrations-sales-pricing-reads-rmq-kube は、RabbitMQ からのメッセージを受け取ると、イベントドリヴンでランタイムを実行します。  
AION の仕様では、Kubernetes 上 の 当該マイクロサービスPod は 立ち上がったまま待機状態で当該メッセージを受け取り、（コンテナ起動などの段取時間をカットして）即座にランタイムを実行します。　

## RabbitMQ への JSON Output

sap-api-integrations-sales-pricing-reads-rmq-kube は、Outputとして、RabbitMQ へのメッセージをJSON形式で出力します。  
Output の サンプルJSON は、Outputs フォルダ内にあります。  

## RabbitMQ の マスタサーバ環境

sap-api-integrations-sales-pricing-reads-rmq-kube が利用する RabbitMQ のマスタサーバ環境は、[rabbitmq-on-kubernetes](https://github.com/latonaio/rabbitmq-on-kubernetes) です。  
当該マスタサーバ環境は、同じエッジコンピューティングデバイスに配置されても、別の物理(仮想)サーバ内に配置されても、どちらでも構いません。

## RabbitMQ の Golang Runtime ライブラリ
sap-api-integrations-sales-pricing-reads-rmq-kube は、RabbitMQ の Golang Runtime ライブラリ として、[rabbitmq-golang-client](https://github.com/latonaio/rabbitmq-golang-client)を利用しています。

## デプロイ・稼働
sap-api-integrations-sales-pricing-reads-rmq-kube の デプロイ・稼働 を行うためには、aion-service-definitions の services.yml に、本レポジトリの services.yml を設定する必要があります。

kubectl apply - f 等で Deployment作成後、以下のコマンドで Pod が正しく生成されていることを確認してください。
```
$ kubectl get pods
```

## 本レポジトリ が 対応する API サービス
sap-api-integrations-sales-pricing-reads-rmq-kube が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_SLSPRCGCONDITIONRECORD_SRV_0001/overview    
* APIサービス名(=baseURL): API_SLSPRICINGCONDITIONRECORD_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-sales-pricing-reads-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* A_SlsPrcgCndnRecdValidity（販売価格条件 - 存在性）※価格条件関連データを取得するために、ToConditionRecord、と合わせて利用されます。
* ToConditionRecord（販売価格条件 - 条件レコード）

## API への 値入力条件 の 初期値
sap-api-integrations-sales-pricing-reads-rmq-kube において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.SalesPricingConditionValidity.Material（品目）
* inoutSDC.SalesPricingConditionValidity.DistributionChannel（流通チャネル）
* inoutSDC.SalesPricingConditionValidity.Customer（得意先）
* inoutSDC.SalesPricingConditionValidity.SalesOrganization（販売組織）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"MaterialDistChannel" が指定されています。    
  
```
	"api_schema": "/sap.s4.beh.salespricingcondition.v1.SalesPricingCondition.Created.v1",
	"accepter": ["MaterialDistChannelCustomer"],
	"condition_record": "",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "/sap.s4.beh.salespricingcondition.v1.SalesPricingCondition.Created.v1",
	"accepter": ["All"],
	"condition_record": "",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetSalesPricingCondition(material, distributionChannel, customer, salesOrganization string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "MaterialDistChannel":
			func() {
				c.MaterialDistChannel(material, distributionChannel)
				wg.Done()
			}()
		case "MaterialDistChannelCustomer":
			func() {
				c.MaterialDistChannelCustomer(material, distributionChannel, customer)
				wg.Done()
			}()
		case "MaterialSalesOrgDistChannel":
			func() {
				c.MaterialSalesOrgDistChannel(material, salesOrganization, distributionChannel)
				wg.Done()
			}()
		case "MaterialSalesOrgDistChannelCustomer":
			func() {
				c.MaterialSalesOrgDistChannelCustomer(material, salesOrganization, distributionChannel, customer)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 販売価格 の 得意先流通チャネル が取得された結果の JSON の例です。  
以下の項目のうち、"ConditionRecord" ～ "to_SlsPrcgConditionRecord" は、/SAP_API_Output_Formatter/type.go 内 の Type PricingConditionValidity {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-sales-pricing-reads/SAP_API_Caller/caller.go#L68",
	"function": "sap-api-integrations-sales-pricing-reads/SAP_API_Caller.(*SAPAPICaller).MaterialDistChannel",
	"level": "INFO",
	"message": [
		{
			"ConditionRecord": "0000008431",
			"ConditionValidityEndDate": "2016-09-15T09:00:00+09:00",
			"ConditionValidityStartDate": "2016-09-01T09:00:00+09:00",
			"ConditionApplication": "V",
			"ConditionType": "PPR0",
			"ConditionReleaseStatus": "",
			"SalesDocument": "",
			"SalesDocumentItem": "0",
			"ConditionContract": "",
			"CustomerGroup": "",
			"CustomerPriceGroup": "",
			"MaterialPricingGroup": "",
			"SoldToParty": "",
			"BPForSoldToParty": "",
			"Customer": "",
			"BPForCustomer": "",
			"PayerParty": "",
			"BPForPayerParty": "",
			"ShipToParty": "",
			"BPForShipToParty": "",
			"Supplier": "",
			"BPForSupplier": "",
			"MaterialGroup": "",
			"Material": "MZ-FG-E15",
			"PriceListType": "",
			"CustomerTaxClassification1": "",
			"ProductTaxClassification1": "",
			"SDDocument": "",
			"ReferenceSDDocument": "",
			"ReferenceSDDocumentItem": "0",
			"SalesOffice": "",
			"SalesGroup": "",
			"SalesOrganization": "1710",
			"DistributionChannel": "10",
			"TransactionCurrency": "",
			"ConditionProcessingStatus": "",
			"PricingDate": "",
			"ConditionScaleBasisValue": "0",
			"TaxCode": "",
			"ServiceDocument": "",
			"ServiceDocumentItem": "0",
			"CustomerConditionGroup": "",
			"to_SlsPrcgConditionRecord": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SLSPRICINGCONDITIONRECORD_SRV/A_SlsPrcgCndnRecdValidity(ConditionRecord='0000008431',ConditionValidityEndDate=datetime'2016-09-15T00%3A00%3A00')/to_SlsPrcgConditionRecord"
		},
		{
			"ConditionRecord": "0000008454",
			"ConditionValidityEndDate": "2017-09-06T09:00:00+09:00",
			"ConditionValidityStartDate": "2016-09-16T09:00:00+09:00",
			"ConditionApplication": "V",
			"ConditionType": "PPR0",
			"ConditionReleaseStatus": "",
			"SalesDocument": "",
			"SalesDocumentItem": "0",
			"ConditionContract": "",
			"CustomerGroup": "",
			"CustomerPriceGroup": "",
			"MaterialPricingGroup": "",
			"SoldToParty": "",
			"BPForSoldToParty": "",
			"Customer": "",
			"BPForCustomer": "",
			"PayerParty": "",
			"BPForPayerParty": "",
			"ShipToParty": "",
			"BPForShipToParty": "",
			"Supplier": "",
			"BPForSupplier": "",
			"MaterialGroup": "",
			"Material": "MZ-FG-E15",
			"PriceListType": "",
			"CustomerTaxClassification1": "",
			"ProductTaxClassification1": "",
			"SDDocument": "",
			"ReferenceSDDocument": "",
			"ReferenceSDDocumentItem": "0",
			"SalesOffice": "",
			"SalesGroup": "",
			"SalesOrganization": "1710",
			"DistributionChannel": "10",
			"TransactionCurrency": "",
			"ConditionProcessingStatus": "",
			"PricingDate": "",
			"ConditionScaleBasisValue": "0",
			"TaxCode": "",
			"ServiceDocument": "",
			"ServiceDocumentItem": "0",
			"CustomerConditionGroup": "",
			"to_SlsPrcgConditionRecord": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SLSPRICINGCONDITIONRECORD_SRV/A_SlsPrcgCndnRecdValidity(ConditionRecord='0000008454',ConditionValidityEndDate=datetime'2017-09-06T00%3A00%3A00')/to_SlsPrcgConditionRecord"
		},
		{
			"ConditionRecord": "0000008454",
			"ConditionValidityEndDate": "9999-12-31T09:00:00+09:00",
			"ConditionValidityStartDate": "2017-11-16T09:00:00+09:00",
			"ConditionApplication": "V",
			"ConditionType": "PPR0",
			"ConditionReleaseStatus": "",
			"SalesDocument": "",
			"SalesDocumentItem": "0",
			"ConditionContract": "",
			"CustomerGroup": "",
			"CustomerPriceGroup": "",
			"MaterialPricingGroup": "",
			"SoldToParty": "",
			"BPForSoldToParty": "",
			"Customer": "",
			"BPForCustomer": "",
			"PayerParty": "",
			"BPForPayerParty": "",
			"ShipToParty": "",
			"BPForShipToParty": "",
			"Supplier": "",
			"BPForSupplier": "",
			"MaterialGroup": "",
			"Material": "MZ-FG-E15",
			"PriceListType": "",
			"CustomerTaxClassification1": "",
			"ProductTaxClassification1": "",
			"SDDocument": "",
			"ReferenceSDDocument": "",
			"ReferenceSDDocumentItem": "0",
			"SalesOffice": "",
			"SalesGroup": "",
			"SalesOrganization": "1710",
			"DistributionChannel": "10",
			"TransactionCurrency": "",
			"ConditionProcessingStatus": "",
			"PricingDate": "",
			"ConditionScaleBasisValue": "0",
			"TaxCode": "",
			"ServiceDocument": "",
			"ServiceDocumentItem": "0",
			"CustomerConditionGroup": "",
			"to_SlsPrcgConditionRecord": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SLSPRICINGCONDITIONRECORD_SRV/A_SlsPrcgCndnRecdValidity(ConditionRecord='0000008454',ConditionValidityEndDate=datetime'9999-12-31T00%3A00%3A00')/to_SlsPrcgConditionRecord"
		},
		{
			"ConditionRecord": "0000009989",
			"ConditionValidityEndDate": "2017-11-15T09:00:00+09:00",
			"ConditionValidityStartDate": "2017-09-07T09:00:00+09:00",
			"ConditionApplication": "V",
			"ConditionType": "PPR0",
			"ConditionReleaseStatus": "",
			"SalesDocument": "",
			"SalesDocumentItem": "0",
			"ConditionContract": "",
			"CustomerGroup": "",
			"CustomerPriceGroup": "",
			"MaterialPricingGroup": "",
			"SoldToParty": "",
			"BPForSoldToParty": "",
			"Customer": "",
			"BPForCustomer": "",
			"PayerParty": "",
			"BPForPayerParty": "",
			"ShipToParty": "",
			"BPForShipToParty": "",
			"Supplier": "",
			"BPForSupplier": "",
			"MaterialGroup": "",
			"Material": "MZ-FG-E15",
			"PriceListType": "",
			"CustomerTaxClassification1": "",
			"ProductTaxClassification1": "",
			"SDDocument": "",
			"ReferenceSDDocument": "",
			"ReferenceSDDocumentItem": "0",
			"SalesOffice": "",
			"SalesGroup": "",
			"SalesOrganization": "1710",
			"DistributionChannel": "10",
			"TransactionCurrency": "",
			"ConditionProcessingStatus": "",
			"PricingDate": "",
			"ConditionScaleBasisValue": "0",
			"TaxCode": "",
			"ServiceDocument": "",
			"ServiceDocumentItem": "0",
			"CustomerConditionGroup": "",
			"to_SlsPrcgConditionRecord": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SLSPRICINGCONDITIONRECORD_SRV/A_SlsPrcgCndnRecdValidity(ConditionRecord='0000009989',ConditionValidityEndDate=datetime'2017-11-15T00%3A00%3A00')/to_SlsPrcgConditionRecord"
		}
	],
	"time": "2022-01-27T22:08:31+09:00"
}
```
