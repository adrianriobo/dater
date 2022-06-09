package xunit

import "testing"

var xunitContentFail string = `
	<testsuites name="crc" tests="41" skipped="0" failures="5" errors="0" time="2723.0541426">
		<testsuite name="Basic test" tests="16" skipped="0" failures="0" errors="0" time="882.3437532">
			<testcase name="CRC clean-up" status="passed" time="4.3474074"/>
		</testsuite>
	</testsuites>
`

var xunitContentSuccess string = `
	<testsuites name="crc" tests="41" skipped="0" failures="0" errors="0" time="2723.0541426">
		<testsuite name="Basic test" tests="16" skipped="0" failures="0" errors="0" time="882.3437532">
			<testcase name="CRC clean-up" status="passed" time="4.3474074"/>
		</testsuite>
	</testsuites>
`

func TestGlobalStatusFail(t *testing.T) {
	status, err := GlobalStatus([]byte(xunitContentFail))
	if status != statusFail || err != nil {
		t.Fatal("Global status should result in Fail")
	}
}

func TestGlobalStatusSuccess(t *testing.T) {
	status, err := GlobalStatus([]byte(xunitContentSuccess))
	if status != statusSuccess || err != nil {
		t.Fatal("Global status should result in Success")
	}
}
