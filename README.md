# sap-api-integrations-product-reads-c4   
sap-api-integrations-product-reads-c4  は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API 製品データを取得するマイクロサービスです。  
sap-api-integrations-product-reads-c4  には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-product-reads-c4  は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/product/overview  

## 動作環境
sap-api-integrations-product-reads-c4  は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-product-reads-c4  は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-product-reads-c4  が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/product/overview 
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-product-reads-c4  には、次の API をコールするためのリソースが含まれています。  

* CampaignCollection（製品 - 製品）※製品の詳細データを取得するために、ProductOtherDescriptions、ProductSalesProcessInformationと合わせて利用されます。
* ProductOtherDescriptions（製品 - 製品外部説明 ※To）
* ProductSalesProcessInformation（製品 - 製品販売プロセス情報 ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-product-reads-c4  において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ProductCollection.ObjectID（対象ID）
* inoutSDC.ProductCollection.ProductID（製品ID）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"ProductCollection" が指定されています。    
  
```
	"api_schema": "Product",
	"accepter": ["ProductCollection"],
	"product_code": "P140100",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "Product",
	"accepter": ["All"],
	"product_code": "P140100",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetProduct(objectID, productID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ProductCollection":
			func() {
				c.ProductCollection(objectID, productID)
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
以下の sample.json の例は、SAP 製品  の 製品データ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "ETag" は、/SAP_API_Output_Formatter/type.go 内 の Type ProductCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona5/bitbucket/sap-api-integrations-product-reads-c4/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-product-reads-c4/SAP_API_Caller.(*SAPAPICaller).ProductCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E03A0701EE288BE9895233EBD27",
			"ProductID": "P140100",
			"UUID": "00163E03-A070-1EE2-88BE-9895233EBD27",
			"Language": "EN",
			"LanguageText": "English",
			"Description": "GS Marengo Womens Mountain Bike",
			"ProductCategoryID": "CUSTOMER-01",
			"Status": "3",
			"StatusText": "Blocked",
			"Usage": "",
			"UsageText": "",
			"Division": "",
			"DivisionText": "",
			"BaseUOM": "EA",
			"BaseUOMText": "Each",
			"CreatedBy": "SAP WORKER",
			"LastChangedBy": "Eddie Smoke",
			"CreatedOn": "2012-10-30T03:16:40+09:00",
			"LastChangedOn": "2015-02-04T06:07:03+09:00",
			"CreatedByUUID": "00163E03-A070-1EE2-88B6-F539A6B028F3",
			"LastChangedByUUID": "00163E03-A070-1EE2-88BA-39BD20F290B5",
			"ExternalSystem": "CRM",
			"ExternalID": "P140100",
			"EntityLastChangedOn": "2015-02-04T06:07:03+09:00",
			"ETag": "2016-07-13T20:41:32+09:00",
			"ProductOtherDescriptions": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/ProductCollection('00163E03A0701EE288BE9895233EBD27')/ProductOtherDescriptions",
			"ProductSalesProcessInformation": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/ProductCollection('00163E03A0701EE288BE9895233EBD27')/ProductSalesProcessInformation"
		}
	],
	"time": "2022-08-29T20:57:02+09:00"
}
```