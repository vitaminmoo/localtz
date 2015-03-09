# localtz

Simple command line utility to convert a time to the local timezone:


```bash
localtz '2015/03/08 23:14:36.576524'
From UTC: 2015/03/08 23:14:36.576524     To Local: 2015/03/08 16:14:36


# Contrived example to show it detects the time's layout:
$ localtz "$(date --utc)"
From UTC: Mon Mar  9 22:25:27 UTC 2015     To Local: Mon Mar  9 15:25:27 PDT 2015


# How to use a different local timezone:
$ localtz -l='America/Chicago' 'Mon Mar  9 22:01:59 UTC 2015'
From UTC: Mon Mar  9 22:01:59 UTC 2015     To America/Chicago: Mon Mar  9 17:01:59 CDT 2015


# Uses timezone from input if it's present:
$ localtz 'Mon Mar  9 17:01:59 CDT 2015'
From CDT: Mon Mar  9 17:01:59 CDT 2015     To Local: Mon Mar  9 10:01:59 PDT 2015
```
