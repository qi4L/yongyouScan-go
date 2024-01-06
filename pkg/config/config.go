package config

import (
	"github.com/gookit/color"
	"sync"
	"yongyouScan/pkg/POC/ConfigurationNc"
	ERP_NC_MLBL "yongyouScan/pkg/POC/ERP-NC-MLBL"
	FileReceiveServlet_Deser "yongyouScan/pkg/POC/FileReceiveServlet-Deser"
	GRP_U8_Proxy_sqljin_xxe "yongyouScan/pkg/POC/GRP-U8-Proxy-sqljin-xxe"
	GRP_U8_U8AppProxy "yongyouScan/pkg/POC/GRP-U8-U8AppProxy"
	GRP_U8_UploadFileData "yongyouScan/pkg/POC/GRP-U8-UploadFileData"
	KSOA_ImageUpload "yongyouScan/pkg/POC/KSOA-ImageUpload"
	KSOA_sqljni "yongyouScan/pkg/POC/KSOA-sqljni"
	MessageServlet_Deser "yongyouScan/pkg/POC/MessageServlet-Deser"
	NC_BshServlet "yongyouScan/pkg/POC/NC-BshServlet"
	u8_MxServlet "yongyouScan/pkg/POC/NC-Cloud-MxServlet"
	NC_JiuQiClientReqDispatch "yongyouScan/pkg/POC/NC-JiuQiClientReqDispatch"
	NC_XbrlPersistenceServlet "yongyouScan/pkg/POC/NC-XbrlPersistenceServlet"
	NC6_5_UploadFile "yongyouScan/pkg/POC/NC6.5-UploadFile"
	NCCloud_FS_sqljni "yongyouScan/pkg/POC/NCCloud-FS-sqljni"
	T_CRM_sqljni "yongyouScan/pkg/POC/T-CRM-sqljni"
	T_DownloadProxy_catfile "yongyouScan/pkg/POC/T-DownloadProxy-catfile"
	T_RecoverPassword "yongyouScan/pkg/POC/T-RecoverPassword"
	T_Uploadfile "yongyouScan/pkg/POC/T-Uploadfile"
	U8_ActionHandlerServlet "yongyouScan/pkg/POC/U8-ActionHandlerServlet"
	U8_CacheInvokeServlet "yongyouScan/pkg/POC/U8-CacheInvokeServlet"
	U8_ClientRequestDispatch "yongyouScan/pkg/POC/U8-ClientRequestDispatch"
	u8_FileTransportServlet "yongyouScan/pkg/POC/U8-FileTransportServlet"
	U8_OA_getSessionList "yongyouScan/pkg/POC/U8-OA-getSessionList"
	U8_OA_test_sqjni "yongyouScan/pkg/POC/U8-OA-test-sqjni"
	U8_RegisterServlet "yongyouScan/pkg/POC/U8-RegisterServlet"
	U8_TaskTreeQuery "yongyouScan/pkg/POC/U8-TaskTreeQuery"
	Uapjs_JNDI "yongyouScan/pkg/POC/Uapjs-JNDI"
	UploadServlet_Deser "yongyouScan/pkg/POC/UploadServlet-Deser"
	accept_upload "yongyouScan/pkg/POC/accept-upload"
	files_Deser "yongyouScan/pkg/POC/files-Deser"
	fs_dlbypass "yongyouScan/pkg/POC/fs-dlbypass"
	fs_rzBypass "yongyouScan/pkg/POC/fs-rzBypass"
	monitorservlet_Desera "yongyouScan/pkg/POC/monitorservlet-Desera"
	u8_LoggingConfigServlet "yongyouScan/pkg/POC/u8-LoggingConfigServlet"
	u8_LoginServlet "yongyouScan/pkg/POC/u8-LoginServlet"
	u8_MonitorServlet "yongyouScan/pkg/POC/u8-MonitorServlet"
	u8_ServletCommander "yongyouScan/pkg/POC/u8-ServletCommander"
	u8_TableInputOperServlet "yongyouScan/pkg/POC/u8-TableInputOperServlet"
	uapws_acessBypass "yongyouScan/pkg/POC/uapws-acessBypass"
	uapws_wsdl_XXE "yongyouScan/pkg/POC/uapws-wsdl-XXE"
)

type WorkExp struct {
	Url string
}

func (c *WorkExp) YonYouScanRun() {
	// 上传写了个一半记得改改
	// color.Blue.Println("[+] 上传的检测可能是不准确的，因为即使是未授权，在请求包中也要包含Cookie")
	color.Blue.Println("[+] URl: " + c.Url)
	var wg sync.WaitGroup
	wg.Add(41)
	go func() {
		ERP_NC_MLBL.Run(c.Url)
		wg.Done()
	}()
	go func() {
		NC_BshServlet.Run(c.Url)
		wg.Done()
	}()
	go func() {
		NCCloud_FS_sqljni.Run(c.Url)
		wg.Done()
	}()
	go func() {
		NC6_5_UploadFile.Run(c.Url)
		wg.Done()
	}()
	go func() {
		NC_XbrlPersistenceServlet.Run(c.Url)
		wg.Done()
	}()
	go func() {
		U8_OA_getSessionList.Run(c.Url)
		wg.Done()
	}()
	go func() {
		U8_OA_test_sqjni.Run(c.Url)
		wg.Done()
	}()
	go func() {
		GRP_U8_UploadFileData.Run(c.Url)
		wg.Done()
	}()
	go func() {
		GRP_U8_Proxy_sqljin_xxe.Run(c.Url)
		wg.Done()
	}()
	go func() {
		Uapjs_JNDI.Run(c.Url)
		wg.Done()
	}()
	go func() {
		T_CRM_sqljni.Run(c.Url)
		wg.Done()
	}()
	go func() {
		T_Uploadfile.Run(c.Url)
		wg.Done()
	}()
	go func() {
		T_RecoverPassword.Run(c.Url)
		wg.Done()
	}()
	go func() {
		GRP_U8_U8AppProxy.Run(c.Url)
		wg.Done()
	}()
	go func() {
		uapws_acessBypass.Run(c.Url)
		wg.Done()
	}()
	go func() {
		fs_rzBypass.Run(c.Url)
		wg.Done()
	}()
	go func() {
		fs_dlbypass.Run(c.Url)
		wg.Done()
	}()
	go func() {
		files_Deser.Run(c.Url)
		wg.Done()
	}()
	go func() {
		T_DownloadProxy_catfile.Run(c.Url)
		wg.Done()
	}()
	go func() {
		KSOA_ImageUpload.Run(c.Url)
		wg.Done()
	}()
	go func() {
		accept_upload.Run(c.Url)
		wg.Done()
	}()
	// https://github.com/wgpsec/YongYouNcTool/blob/main/src/main/java/toolPannel.java
	go func() {
		MessageServlet_Deser.Run(c.Url)
		wg.Done()
	}()
	go func() {
		UploadServlet_Deser.Run(c.Url)
		wg.Done()
	}()
	go func() {
		monitorservlet_Desera.Run(c.Url)
		wg.Done()
	}()
	go func() {
		FileReceiveServlet_Deser.Run(c.Url)
		wg.Done()
	}()
	// https://blog.csdn.net/qq_41904294/article/details/134908263
	go func() {
		u8_TableInputOperServlet.Run(c.Url)
		wg.Done()
	}()
	go func() {
		u8_LoginServlet.Run(c.Url)
		wg.Done()
	}()
	go func() {
		u8_FileTransportServlet.Run(c.Url)
		wg.Done()
	}()
	go func() {
		U8_CacheInvokeServlet.Run(c.Url)
		wg.Done()
	}()
	go func() {
		U8_ActionHandlerServlet.Run(c.Url)
		wg.Done()
	}()
	go func() {
		u8_ServletCommander.Run(c.Url)
		wg.Done()
	}()
	go func() {
		u8_MxServlet.Run(c.Url)
		wg.Done()
	}()
	go func() {
		u8_MonitorServlet.Run(c.Url)
		wg.Done()
	}()
	go func() {
		U8_ClientRequestDispatch.Run(c.Url)
		wg.Done()
	}()
	go func() {
		U8_RegisterServlet.Run(c.Url)
		wg.Done()
	}()
	go func() {
		u8_LoggingConfigServlet.Run(c.Url)
		wg.Done()
	}()

	go func() {
		uapws_wsdl_XXE.Run(c.Url)
		wg.Done()
	}()
	go func() {
		ConfigurationNc.Run(c.Url)
		wg.Done()
	}()
	go func() {
		KSOA_sqljni.Run(c.Url)
		wg.Done()
	}()
	go func() {
		NC_JiuQiClientReqDispatch.Run(c.Url)
		wg.Done()
	}()
	go func() {
		U8_TaskTreeQuery.Run(c.Url)
		wg.Done()
	}()
	wg.Wait()
}
