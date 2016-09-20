#!/bin/bash

manager_host="10.117.7.234"
passwd="vmware"

RestartService() {
	echo "Clean exist containers...\n"

	containers=$(docker ps -a | awk '{print $1}'| grep -v CONTAINER)
	if [ ! -n "$containers" ]
	then
		echo "There is no contaners in this host!"
	else
		echo "rm the exist containers!"
		docker rm -f $containers 
	fi

	echo "\nRestart docker service and iptables firewall...\n"

	sudo service docker stop
	sudo service ufw restart
	sudo service docker start

	echo "Service is already started!"
}

DeploySwarmManager() {
	
	echo "\nDeploy consul on the manager node..."
	docker run -d --restart=always -p 8500:8500 --name=consul progrium/consul -server -bootstrap

	echo "\nDeploy swarm manager on the manager node..."
	docker run -d -p 4000:4000 --restart=always --name swarm-manager swarm --experimental manage -H :4000 --replication --advertise ${manager_host}:4000 consul://${manager_host}:8500
}

RestartNodes() {
	restart_service_cmd="docker rm -f \$(docker ps -a | awk '{print \$1}'| grep -v CONTAINER);service docker stop;service ufw restart;service docker start"
	i=1
	for node_ip in $*
	do
		echo "Restart service on node-0${i}:${node_ip}..."
		sshpass -p ${passwd} ssh root@${node_ip}  ${restart_service_cmd}
		i=`expr $i + 1`;
	done

} 

DeploySwarmNodes() {
	i=1
	for node_ip in $*
	do
		echo "Deploy swarm on node-0${i}:${node_ip}..."
		run_swarm_cmd="docker run -d --name swarm-node0${i}  --restart=always  swarm --experimental join --advertise=${node_ip}:2375 consul://${manager_host}:8500"
		sshpass -p ${passwd} ssh work@${node_ip} ${run_swarm_cmd}
		i=`expr $i + 1`;
	done
}


RestartService
RestartNodes "10.117.7.164" "10.117.4.251" "10.117.5.196"
DeploySwarmManager
DeploySwarmNodes "10.117.7.164" "10.117.4.251" "10.117.5.196"

