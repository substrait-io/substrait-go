// SPDX-License-Identifier: Apache-2.0

package substraitproto

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func requireProtoOneof(message protoreflect.Message, name protoreflect.Name) error {
	if !message.IsValid() {
		return fmt.Errorf("missing message for %s", name)
	}
	field := message.WhichOneof(message.Descriptor().Oneofs().ByName(name))
	if field == nil {
		return fmt.Errorf("missing %s", name)
	}
	if field.Message() != nil && !message.Get(field).Message().IsValid() {
		return fmt.Errorf("missing %s payload", field.Name())
	}
	return nil
}
