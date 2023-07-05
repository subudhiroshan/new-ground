#!/usr/bin/env python

import http.client
import json

def req(host, method):
    if host == "backend":
        conn = http.client.HTTPConnection("localhost", 8080)
    if host == "frontend":
        conn = http.client.HTTPConnection("localhost", 8090)
    conn.request("GET", method)
    return conn.getresponse().read().decode().strip()

class E2ETest:
    def __init__(self, name, setup, check):
        self.name = name
        self.setup = setup
        self.check = check

    def run(self):
        for s in self.setup:
            host, method = s
            req(host, method)

        host, method, expected = self.check
        got = req(host, method)

        if got == expected:
            return (True, "")

        return (False, f"Expected \"{expected}\", Got \"{got}\"")

tests = [
        E2ETest(
            "smoke",
            [
                ("backend", "/"),
                ("frontend", "/"),
            ],
            ("frontend", "", "0"),
            ),
        E2ETest(
            "negative-start",
            [
                ("backend", "/decrease"),
                ("backend", "/decrease"),
                ("backend", "/decrease"),
                ("frontend", "/process"),
            ],
            ("frontend", "", "-6"),
            ),
        E2ETest(
            "increase-even",
            [
                ("backend", "/increase"),
                ("backend", "/increase"),
                ("frontend", "/process"),
            ],
            ("frontend", "", "3"),
            ),
        ]

passed = 0
failed = 0
test_runs = []

for t in tests:
    req("backend", "/zero")
    success, reason = t.run()
    if success:
        passed += 1
        test_runs.append(dict(name=t.name, success=success, reason=""))
    else:
        failed += 1
        test_runs.append(dict(name=t.name, success=success, reason=reason))


results = dict(
        summary=dict(
            passed=passed,
            failed=failed,
            total=passed+failed,
            ),
        tests=test_runs,
        )
print(json.dumps(results, indent=4))
