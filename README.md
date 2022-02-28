# pingtest
ping test tool designed for a specific purpose (error codes on failure rates) for crontab monitoring

## Usage

```
% ./pingtest.darwin.amd64 
Usage of ./pingtest.darwin.amd64:
  -host string
        hostname to ping
  -interval duration
        how long to wait between ping attempts; also defines each ping and total timeout (default 1s)
  -number int
        number of pings to run (default 1)
  -privileged
        run in privileged mode
  -quiet
        do not output any logging except for fatal errors and final result
  -rtts
        set to print each individual packet RTT, not just stats; setting quiet disables this too
  -threshold int
        packet loss percentage threshold, when test is considered to be failed; will return with errno 127 (default 100)
```

## Example reboot on too many failures

If failure rate 50% with 15 packets, 4 seconds apart means, more or less, if connectivity is not there for half the time over one minute. Monitor connection to local router.

```
% sudo ./pingtest.linux.amd64 -host 192.168.0.1 -interval 4s -number 15  -privileged -quiet -threshold 50 || reboot
```

## Example reboot if connectivity not there for a minute, test 4 times

If failure rate 100% with 4 packets, 15 seconds apart, fairly certain that connectivity is lost for over a minute, or interface is flapping constantly.

```
% sudo ./pingtest.linux.amd64 -host 192.168.0.1 -interval 15s -number 4  -privileged -quiet -threshold 100 || reboot
```

## Example reboot if connectivity not there for 2 minutes, test 4 times/minute

If failure rate 100% with 4 packets, 15 seconds apart, fairly certain that connectivity is lost for over a minute, or interface is flapping constantly.

```
% sudo ./pingtest.linux.amd64 -host 192.168.0.1 -interval 15s -number 8  -privileged -quiet -threshold 100 || reboot
```

## Example reboot if connectivity not there for 2 minutes, test 4 times/minute

If failure rate 87% with 4 packets, 15 seconds apart, fairly certain that connectivity is lost for over a minute, or interface is flapping constantly.

In this case, if `7/8*100=87.5%`. This means, if one out of 8 packets sent over 2 minutes gets through, we will still reboot (over 87% packet loss).

```
% sudo ./pingtest.linux.amd64 -host 192.168.0.1 -interval 15s -number 8  -privileged -quiet -threshold 87 || reboot
```
