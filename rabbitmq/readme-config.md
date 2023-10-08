
1. Add the RabbitMQ APT repository to your system:
   sudo apt install -y curl gnupg
   curl -fsSL https://packages.erlang-solutions.com/ubuntu/erlang_solutions.asc | sudo gpg --dearmor -o /usr/share/keyrings/erlang-solutions-archive-keyring.gpg
2. Add the RabbitMQ APT repository to your system:
   echo "deb [signed-by=/usr/share/keyrings/erlang-solutions-archive-keyring.gpg] https://packages.erlang-solutions.com/ubuntu $(lsb_release -cs) contrib" | sudo tee /etc/apt/sources.list.d/rabbitmq.list
3. Install RabbitMQ and its dependencies:
   sudo apt update
   sudo apt install -y erlang-base erlang-asn1 erlang-crypto erlang-eldap erlang-ftp erlang-inets erlang-mnesia erlang-os-mon erlang-parsetools erlang-public-key erlang-runtime-tools erlang-snmp erlang-ssl erlang-syntax-tools erlang-tftp erlang-tools erlang-xmerl rabbitmq-server
4. Start and Enable RabbitMQ Service
   sudo systemctl start rabbitmq-server
   sudo systemctl enable rabbitmq-server
5. Create a new RabbitMQ user (replace username and password with your preferred values):
   sudo rabbitmqctl add_user biangacila Nathan010309*
6. Give the new user administrative privileges:
   sudo rabbitmqctl set_user_tags biangacila administrator
   sudo rabbitmqctl set_permissions -p / biangacila ".*" ".*" ".*"
7. Adjust Firewall Rules
   sudo ufw allow 5672/tcp
   sudo ufw allow 15672/tcp
8. Access RabbitMQ Management UI (Optional)
   To access the RabbitMQ Management UI from your local machine, you can create an SSH tunnel:
   ssh -L 15672:localhost:15672 username@your_server_ip
9. Then, open a web browser and go to http://localhost:15672 to access the RabbitMQ Management UI. Log in with the credentials you created earlier.





