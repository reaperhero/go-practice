package regexp

import (
	"fmt"
	"regexp"
	"testing"
)

func TestEmail(t *testing.T) {
	text := "My email is ccmouse@gmail.com"
	compile := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := compile.FindString(text)
	fmt.Println(match)
}

func TestImage(t *testing.T) {
	text := `- /bin/sh
        - -c
        - |-
          res=ls /plugin/1.10_release_5.1.x_862 | wc -l
          if [ -e "/plugin/1.10_release_5.1.x_862" ] && [ $res != 0 ]
          then
            echo "plugin already exists"
          else
            echo "plugin has changed, copy plugin from image"
            sshpass -p 'Abc!@#135' ssh -o StrictHostKeyChecking=no root@172.16.82.40 "docker run --rm -t -v /data/nfsfile/dtstack-plugin/base53/default-53x-new-easyindex/sqlplugin/1.10_release_5.1.x_862:/plugintmp 172.16.84.106/dtstack-relea se/sqlplugin:1.10_release_5.1.x_862 172.16.84.106/dtstack-release/sqlplugin:1.10_release_5.1.x_862 /bin/sh -c \"cp -r /data/insight_plugin//* /plugintmp/\""
            if [ $? != 0 ]; then
              echo "copy plugin failed"
              rm -rf /plugin/1.10_release_5.1.x_862
              exit 1
            fi
            cd /plugin
            ls /plugin | grep -v 1.10_release_5.1.x_862 | xargs rm -rf
            echo "plugin is ready"
          fi`
	compile := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+/[^ ]+:[\.|\w\_|\-\_]+`)
	match := compile.FindStringSubmatch(text)
	for _, s := range match {
		fmt.Println(s)
	}
}

func TestName(t *testing.T) {
	data := "      uic_server_name: ${tengine.uic_server_name} # config_type: availability; comment: 用户; explain: 配置秒数详细信息"
	configTypeReg := regexp.MustCompile("config_type:(?s:(.*?));")
	submatch := configTypeReg.FindStringSubmatch(data)
	fmt.Println(submatch[1])
}
