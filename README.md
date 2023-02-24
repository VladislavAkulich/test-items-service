# test-items-service

## How to run DB local

1. Install Makefile support (Windows tips here, run through the cmd or powershell https://linuxhint.com/run-makefile-windows/)
2. Install Docker hub
3. Execute command in terminal `make db_up` to start DB container, `make db_stop` to stop.
4. For now u also need to apply migrations manually from ide. (see `migrations/` folder)