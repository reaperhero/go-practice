package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/clbanning/mxj"
	"testing"
)

var defaultCapacityScheduler = `<!--
  Licensed under the Apache License, Version 2.0 (the "License");
-->
<configuration>

  <property>
    <name>yarn.scheduler.capacity.maximum-applications</name>
    <value>10000</value>
    <description>
      Maximum number of applications that can be pending and running.
    </description>
  </property>

  <property>
    <name>yarn.scheduler.capacity.maximum-am-resource-percent</name>
    <value>1</value>
    <description>
      Maximum percent of resources in the cluster which can be used to run
      application masters i.e. controls number of concurrent running
      applications.
    </description>
  </property>

  <property>
    <name>yarn.scheduler.capacity.resource-calculator</name>
    <value>org.apache.hadoop.yarn.util.resource.DominantResourceCalculator</value>
    <description>
      The ResourceCalculator implementation to be used to compare
      Resources in the scheduler.
      The default i.e. DefaultResourceCalculator only uses Memory while
      DominantResourceCalculator uses dominant-resource to compare
      multi-dimensional resources such as Memory, CPU etc.
    </description>
  </property>

  <property>
    <name>yarn.scheduler.capacity.root.queues</name>
    <value>default</value>
    <description>
      The queues at the this level (root is the root queue).
    </description>
  </property>

  <property>
    <name>yarn.scheduler.capacity.root.default.capacity</name>
    <value>100</value>
    <description>Default queue target capacity.</description>
  </property>

  <property>
    <name>yarn.scheduler.capacity.root.default.user-limit-factor</name>
    <value>1</value>
    <description>
      Default queue user limit a percentage from 0.0 to 1.0.
    </description>
  </property>

  <property>
    <name>yarn.scheduler.capacity.root.default.maximum-capacity</name>
    <value>100</value>
    <description>
      The maximum capacity of the default queue.
    </description>
  </property>

  <property>
    <name>yarn.scheduler.capacity.root.default.state</name>
    <value>RUNNING</value>
    <description>
      The state of the default queue. State can be one of RUNNING or STOPPED.
    </description>
  </property>

  <property>
    <name>yarn.scheduler.capacity.root.default.acl_submit_applications</name>
    <value>*</value>
    <description>
      The ACL of who can submit jobs to the default queue.
    </description>
  </property>

  <property>
    <name>yarn.scheduler.capacity.root.default.acl_administer_queue</name>
    <value>*</value>
    <description>
      The ACL of who can administer jobs on the default queue.
    </description>
  </property>

  <property>
    <name>yarn.scheduler.capacity.node-locality-delay</name>
    <value>40</value>
    <description>
      Number of missed scheduling opportunities after which the CapacityScheduler
      attempts to schedule rack-local containers.
      Typically this should be set to number of nodes in the cluster, By default is setting
      approximately number of nodes in one rack which is 40.
    </description>
  </property>

  <property>
    <name>yarn.scheduler.capacity.queue-mappings</name>
    <value></value>
    <description>
      A list of mappings that will be used to assign jobs to queues
      The syntax for this list is [u|g]:[name]:[queue_name][,next mapping]*
      Typically this list will be used to map users to queues,
      for example, u:%user:%user maps all users to queues with the same name
      as the user.
    </description>
  </property>

  <property>
    <name>yarn.scheduler.capacity.queue-mappings-override.enable</name>
    <value>false</value>
    <description>
      If a queue mapping is present, will it override the value specified
      by the user? This can be used by administrators to place jobs in queues
      that are different than the one specified by the user.
      The default is false.
    </description>
  </property>


</configuration>`

type Property struct {
	XMLName     xml.Name `xml:"property"`
	Text        string   `xml:",chardata"`
	Name        string   `xml:"name"`
	Value       string   `xml:"value"`
	Description string   `xml:"description"`
}

// 	ss, sserr = m.ValuesForKey("book", "author:William T. Gaddis")
// 	ss, sserr = m.ValuesForKey("author", "-seq:1")
// 	ss, sserr := m.ValuesForPath("doc.books.book.author")
// 	ss, sserr = m.ValuesForPath("doc.books.*", "-seq:3")
// 	ss, sserr = m.ValuesForPath("doc.*.*", "-seq:3")
// 	ss, sserr := m.ValuesForPath("doc.books.book", "!author:William T. Gaddis")
// 	v, verr := m.ValuesForPath("section1.data[0].F1")
// 	p, _ := m.UpdateValuesForPath("author:John Hawkes", "doc.books.book", "title:The Beetle Leg")
// 	n, _ = m.UpdateValuesForPath("author:William T. Gaddis", "doc.books.book.*", "title:The Recognitions")
//  mv.Exists("Div.Color")
// 	err := mv.SetValueForPath("big", "Div.Font.Size")
//  mv.Remove("Div.Colour")
func TestName(t *testing.T) {
	capSch, _, err := mxj.NewMapXmlReaderRaw(bytes.NewBuffer([]byte(defaultCapacityScheduler)))
	if err != nil {
		panic(err)
	}
	// 查询
	valueForKey, err := capSch.ValuesForPath("configuration.property", "name:yarn.scheduler.capacity.root.queues")
	if err != nil {
		panic(err)
	}
	fmt.Println(valueForKey)

	// 修改,会同步修改上面的valueForKey，因为都是map引用
	_, err = capSch.UpdateValuesForPath("value:default111", "configuration.property", "name:yarn.scheduler.capacity.root.queues")
	if err != nil {
		panic(err)
	}

	// 删除
	ps, err := capSch.ValuesForPath("configuration.property")
	if err != nil {
		panic(err)
	}
	for i, p := range ps {
		switch p.(type) {
		case map[string]interface{}:
			if pm, ok := p.(map[string]interface{}); ok && pm["name"].(string) == "yarn.scheduler.capacity.root.queues" {
				ps = append(ps[:i], ps[i+1:]...)
			}
		}
	}
	_, err = capSch.UpdateValuesForPath(mxj.Map{"property": ps}, "configuration")
	if err != nil {
		panic(err)
	}
	// 增加
	addProperty, err := capSch.ValuesForPath("configuration.property")
	addProperty = append(addProperty, mxj.Map{
		"name":        "asdas",
		"value":       "sdas",
		"description": "asdsadasd",
	})
	_, err = capSch.UpdateValuesForPath(mxj.Map{"property": addProperty}, "configuration")
	if err != nil {
		panic(err)
	}

	// 打印
	vvv, err := capSch.XmlIndent("  ", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(vvv))
}
