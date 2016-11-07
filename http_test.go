package Figo

import "testing"

func TestDownLoad(t *testing.T) {

	remoteFileName := "http://www.apache.org/dist/tomcat/tomcat-7/v7.0.72/bin/apache-tomcat-7.0.72.zip"
	localFileName := "./apache-tomcat-7.0.72.zip"
	DownLoad(localFileName, remoteFileName, 10)
}
