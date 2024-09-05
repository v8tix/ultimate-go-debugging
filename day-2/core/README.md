When you run a Go program with `GOTRACEBACK=crash`, it should generate a core dump if the program crashes. However, if the core file is not being created, here are some possible reasons and steps to troubleshoot:

1. **Core Dumps Are Disabled or Limited:**
    - On many systems, core dumps are disabled or limited in size by default. You can check the current core dump settings with the `ulimit` command:
      ```bash
      ulimit -c
      ```
      If it returns `0`, core dumps are disabled. You can enable core dumps and set a limit (or make it unlimited) with:
      ```bash
      ulimit -c unlimited
      ```

2. **System Configuration:**
    - On some systems, there are additional configurations that control where core dumps are saved and under what conditions. For example, on Linux, you might need to check `/proc/sys/kernel/core_pattern`:
      ```bash
      cat /proc/sys/kernel/core_pattern
      ```
      This file determines the location and naming pattern for core dumps. If it's set to something unexpected, you may need to modify it to save core dumps in a specific location.

3. **Permission Issues:**
    - Ensure that the directory where the core file is supposed to be created is writable by the user running the program. If the program does not have the necessary permissions, the core file won't be created.

4. **Executable Format:**
    - The core dump will only be generated if the program is compiled and linked in a way that supports core dumps. Ensure that your program is not stripped (`go build` doesn't strip binaries by default, but if you're using additional tools, this could be a factor).

5. **Check Application and Logs:**
    - Verify that the application is indeed crashing in a way that would generate a core dump. You can look at the system logs (`/var/log/syslog`, `/var/log/messages`, etc.) for any relevant error messages.

6. **Running the Command:**
    - When you run the command:
      ```bash
      GOTRACEBACK=crash ./bin/crash
      ```
      Make sure you are running this in a shell that respects environment variable settings. If you're using a shell script, double-check that the `GOTRACEBACK=crash` is being properly applied.

If none of these steps solve the problem, you might want to check if there are any specific settings or restrictions related to core dumps on your particular operating system or environment.

#################

The output you received from `cat /proc/sys/kernel/core_pattern` indicates that core dumps are being handled by Apport, a crash reporting system used in Ubuntu and some other Linux distributions. Apport intercepts core dumps and processes them instead of simply writing them to disk as a core file.

To generate a standard core dump file instead of using Apport, you can change the `core_pattern` to a simple file path, such as `/tmp/core-%e.%p`, where `%e` is the executable name and `%p` is the process ID. Here's how you can do it:

1. **Temporarily Change Core Pattern:**
   You can temporarily change the core pattern to save core dumps to a specific directory, like `/tmp`:
   ```bash
   sudo sysctl -w kernel.core_pattern=/tmp/core-%e.%p
   ```
   This command changes the core dump pattern for the current session. After executing this command, any crashing program will generate a core dump file in the `/tmp` directory with a name like `core-crash.1234` (where `crash` is the executable name and `1234` is the process ID).

2. **Persist the Change:**
   If you want to make this change permanent, you can add it to `/etc/sysctl.conf`:
   ```bash
   echo "kernel.core_pattern=/tmp/core-%e.%p" | sudo tee -a /etc/sysctl.conf
   sudo sysctl -p
   ```

3. **Re-run Your Program:**
   After making this change, run your program again with `GOTRACEBACK=crash`:
   ```bash
   GOTRACEBACK=crash ./bin/crash
   ```
   If the program crashes, it should generate a core dump file in the `/tmp` directory.

4. **Revert to Apport (Optional):**
   If you want to revert to the default behavior with Apport, you can run:
   ```bash
   sudo sysctl -w kernel.core_pattern="|/usr/share/apport/apport -p%p -s%s -c%c -d%d -P%P -u%u -g%g -- %E"
   ```

This should allow you to collect the core dump file for further analysis.

#########

To generate the core dump in the same directory where the binary is located, you can modify the `kernel.core_pattern` to include a relative path or specify the current working directory (`.`) where the binary is executed. Here's how you can do it:

1. **Set Core Pattern to Current Directory:**
   You can set the core pattern to the current directory by updating `kernel.core_pattern` as follows:
   ```bash
   sudo sysctl -w kernel.core_pattern=./core-%e.%p
   ```
   This will create core dumps in the directory from which the binary is executed, with a filename pattern `core-binaryname.pid`.

2. **Persist the Change (Optional):**
   If you want to make this change permanent, add it to `/etc/sysctl.conf`:
   ```bash
   echo "kernel.core_pattern=./core-%e.%p" | sudo tee -a /etc/sysctl.conf
   sudo sysctl -p
   ```

3. **Run Your Program:**
   Now, when you run your program:
   ```bash
   GOTRACEBACK=crash ./bin/crash
   ```
   If the program crashes, it will generate a core dump file in the same directory as `./bin/crash`, with a name like `core-crash.1234`.

### Important Considerations:
- **Write Permissions:** Ensure that the directory where the binary is located has write permissions for the user running the program. Otherwise, the core dump file cannot be created.

- **File Name Length:** Some filesystems have limitations on filename length, so keep the pattern simple to avoid issues.

This configuration will ensure that core dumps are generated directly in the directory where the binary resides.