//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package renders

import (
    "internal/vars"
    "fmt"
    "testing"
)

func TestURLSchemaRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    templatedInput := fmt.Sprintf("%s://127.0.0.1:80/index.html", vars.EutherpeTemplateNeedleURLSchema)
    eutherpeVars.HTTPd.URLSchema = "http"
    output := URLSchemaRender(templatedInput, eutherpeVars)
    if output != "http://127.0.0.1:80/index.html" {
        t.Errorf("URLSchemaRender() seems not to be working accordingly.\n")
    }
}
