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

func TestCopyrightRender(t *testing.T) {
    output := CopyrightRender(vars.EutherpeTemplateNeedleCopyrightDisclaimer,
                              nil)
    if output != vars.EutherpeCopyrightDisclaimer {
        t.Errorf("CopyrightDisclaimer() is not rendering accordingly.\n")
    }
}
