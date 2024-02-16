package seleniumchromescrapper

import (
	"archive/zip"
	"fmt"
	"os"
)

const (
	manifestFName = "manifest.json"
	backFName     = "background.js"
)

// BuildProxyExtension creates a chrome extension for proxy authentication
func BuildProxyExtension(zipFName, host, port, userName, password string) error {
	manifestJSON := `{
		  "version": "1.0.0",
		  "manifest_version": 2,
		  "name": "Chrome Proxy",
		  "permissions": [
			"proxy",
			"tabs",
			"unlimitedStorage",
			"storage",
			"<all_urls>",
			"webRequest",
			"webRequestBlocking"
		  ],
		  "background": {
			"scripts": ["background.js"]
		  },
		  "minimum_chrome_version":"22.0.0"
		}`

	backgroundJS := fmt.Sprintf(`var config = {
		  mode: "fixed_servers",
		  rules: {
			singleProxy: {
			  scheme: "http",
			  host: "%s",
			  port: parseInt(%s)
			},
			bypassList: ["localhost"]
		  }
		};

		chrome.proxy.settings.set({value: config, scope: "regular"}, function() {});

		function callbackFn(details) {
		  return {
			authCredentials: {
			  username: "%s",
			  password: "%s"
			}
		  };
		}

		chrome.webRequest.onAuthRequired.addListener(
		  callbackFn,
		  {urls: ["<all_urls>"]},
		  ['blocking']
		);`, host, port, userName, password)

	fos, err := os.Create(zipFName)
	if err != nil {
		return fmt.Errorf("os.Create Error: %w", err)
	}
	zipWriter := zip.NewWriter(fos)
	defer zipWriter.Close()
	defer fos.Close()

	if err := createFile(zipWriter, manifestFName, []byte(manifestJSON)); err != nil {
		return err
	}

	if err := createFile(zipWriter, backFName, []byte(backgroundJS)); err != nil {
		return err
	}

	// copy file to zip archive
	if err := zipWriter.Close(); err != nil {
		return fmt.Errorf("zipWriter.Close error: %w", err)
	}

	return nil
}

func createFile(zipWriter *zip.Writer, fileName string, content []byte) error {
	file, err := zipWriter.Create(fileName)
	if err != nil {
		return fmt.Errorf("zipWriter error: %w", err)
	}
	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("file.Write error: %w", err)
	}

	return nil
}
