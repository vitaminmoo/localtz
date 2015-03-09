# localtz

Simple command line utility to convert a time to the local timezone:


```bash
localtz '2015/03/08 23:14:36.576524'
# Outputs: 2015/03/08 16:14:36


# Contrived example to show it detects the time's layout:
./localtz "$(date --utc)"
# Outputs: Mon Mar  9 14:54:23 PDT 2015
```
