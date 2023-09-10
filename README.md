# Node IP Passthru to Auto Instrumentation

This is a sample project to show how you can pass Node IP to auto
instrumentation configured using the OpenTelemetry Operator. It assumes you have
a kind cluster (w/ctlptl) running the OpenTelemetry Operator. 

To use, run `tilt up`. Set a Honeycomb key in the `collector.yaml` file to send
traces to Honeycomb.