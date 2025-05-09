package mermaid

/*
 * SPDX-FileCopyrightText: 2024 Siemens AG
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * Author: Michael Adler <michael.adler@siemens.com>
 */

import (
	"bytes"
	"testing"

	"github.com/siemens/wfx/workflow/dau"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	buf := new(bytes.Buffer)
	gen := NewGenerator()
	err := gen.Generate(buf, dau.DirectWorkflow())
	require.NoError(t, err)
	actual := buf.String()
	expected := `stateDiagram-v2
    [*] --> INSTALL
    INSTALL --> INSTALLING: CLIENT
    INSTALL --> TERMINATED: CLIENT
    INSTALLING --> INSTALLING: CLIENT
    INSTALLING --> TERMINATED: CLIENT
    INSTALLING --> INSTALLED: CLIENT
    INSTALLED --> ACTIVATE: WFX
    ACTIVATE --> ACTIVATING: CLIENT
    ACTIVATE --> TERMINATED: CLIENT
    ACTIVATING --> ACTIVATING: CLIENT
    ACTIVATING --> TERMINATED: CLIENT
    ACTIVATING --> ACTIVATED: CLIENT
    ACTIVATED --> [*]
    TERMINATED --> [*]
    classDef cl_INSTALL color:black,fill:#00cc00
    class INSTALL cl_INSTALL
    classDef cl_INSTALLING color:black,fill:#00cc00
    class INSTALLING cl_INSTALLING
    classDef cl_INSTALLED color:black,fill:#00cc00
    class INSTALLED cl_INSTALLED
    classDef cl_ACTIVATE color:black,fill:#00cc00
    class ACTIVATE cl_ACTIVATE
    classDef cl_ACTIVATING color:black,fill:#00cc00
    class ACTIVATING cl_ACTIVATING
    classDef cl_ACTIVATED color:black,fill:#4993dd
    class ACTIVATED cl_ACTIVATED
    classDef cl_TERMINATED color:black,fill:#9393dd
    class TERMINATED cl_TERMINATED
    Note right of INSTALL: <b>Group to Color Mapping</b><br/><font color="#00cc00">OPEN</font> - regular workflow-advancing states<br/><font color="#4993dd">CLOSED</font> - a successful update's terminal states<br/><font color="#9393dd">FAILED</font> - a failed update's terminal states
`
	assert.Equal(t, expected, actual)
}
