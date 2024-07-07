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
    "testing"
)

func TestVersionRender(t *testing.T) {
    output := VersionRender(vars.EutherpeTemplateNeedleVersion, nil)
    if output != vars.EutherpeVersion {
        t.Errorf("VersionRender() is not rendering accordingly.\n")
    }
}
