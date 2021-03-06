package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hashicorp/errwrap"
	"github.com/wvanbergen/kazoo-go"
)

// SanityTestTopicPlanOpts represents the 'sanity-test-topic-plan' command
type SanityTestTopicPlanOpts struct {
}

// Execute is callback from go-flags.Commander interface
func (c SanityTestTopicPlanOpts) Execute(_ []string) (err error) {
	decoder := json.NewDecoder(os.Stdin)
	// creds := Credentials{}
	creds := make(map[string]string)
	err = decoder.Decode(&creds)
	if err != nil {
		return errwrap.Wrapf("Failed to unmarshal credentials: {{err}}", err)
	}
	fmt.Printf("Loaded credentials: %#v\n", creds)

	var topicName = creds["topicName"]
	var zkPeers = creds["zkPeers"]
	var hostname = creds["hostname"]

	if topicName == "" {
		return fmt.Errorf("'topicName' was not provided")
	}
	if zkPeers == "" {
		return fmt.Errorf("'zkPeers' was not provided")
	}
	if hostname == "" {
		return fmt.Errorf("'hostname' was not provided")
	}

	zkConf := kazoo.NewConfig()
	kz, err := kazoo.NewKazooFromConnectionString(zkPeers, zkConf)
	if err != nil {
		return errwrap.Wrapf("Could not connect to Kafka: {{err}}", err)
	}
	defer func() { _ = kz.Close() }()
	exists, err := kz.Topic(topicName).Exists()
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Topic %s could not be looked up: {{err}}", topicName), err)
	}
	if !exists {
		return fmt.Errorf("Topic %s does not exist", topicName)
	}

	return
}
