# Cron Scheduler

### Purpose
Written as a Technical test for [Lucid](https://luc.id/) for the application of Software Engineer.

### The Problem
We have a set of tasks, each running at least daily, which are scheduled with a simplified cron. We want to find when each of them will next run.

The scheduler configuration looks like this:

```text
30 1 /bin/run_me_daily

45 * /bin/run_me_hourly

* * /bin/run_me_every_minute

* 19 /bin/run_me_sixty_times
```

The first field is the minutes past the hour, the second field is the hour of the day and the third is the command to run. For both cases `*` means that it should run for all values of that field. In the above example, `run_me_daily` has been set to run at 1:30 am every day and `run_me_hourly` at 45 minutes past the hour every hour. The fields
 are whitespace separated and each entry is on a separate line.

We would like you to write a command line program that outputs the soonest time at which each of the commands will run and whether this will be `today` or `tomorrow`.

The config input will be via STDIN, and the output should be via STDOUT. The `current time` will be the single command line argument to the program in the format `HH:MM`.

For example, given the above examples as input and the simulated `current time` command-line argument `16:10` the output should be

```text
1:30 tomorrow - /bin/run_me_daily

16:45 today - /bin/run_me_hourly

16:10 today - /bin/run_me_every_minute

19:00 today - /bin/run_me_sixty_times
```

### Usage
When running the application the first command-line argument represents the `current-time`. That is, the time that the application should compare the cron jobs to. If this argument is omitted, the application will default the `current-time` parameter to the time the application is run. 

When running the application, you can use STDIN to input your cron jobs, delimited by a newline. Each line should have comprise of 3 elements (Hour, Minute and Cron Job). If a line is not valid, it will not be processed, but the application will continue.

Once you are finished, send a `end` command to the application and it will process your request. 

- To run the application in development mode, use the `run.sh` script. 

### Tests
All tests can be run using the `tests.sh` script.