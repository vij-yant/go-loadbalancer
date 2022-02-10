# Loadbalancer-in-go

This project is an implementation of a loadbalancer from scratch using a load balancing algorithm called round-robin.Health-checker is implemented as well to validate the server's availability periodically.

#### What's a Load Balancer?
![1_tEaZGz-p1-E2ytNjl5RPJg](https://user-images.githubusercontent.com/79076537/153355762-9124db33-f15b-4aaa-a463-56e73d564e4f.jpeg)

A load balancer acts as a ‘reverse-proxy’ to represent the application servers to the client through a virtual IP address (VIP).
Load balancers are used to provide availability and scalability to the application. The application can scale beyond the capacity of a single server.
The load balancer works to steer the traffic to a pool of available servers through various otherload balancing algorithms like Least Connection, Weighted Response Time etc.

#### How it works?

A load balancer is a reverse proxy. It presents a virtual IP address (VIP) representing the application to the client. The client connects to the VIP and the load balancer makes a determination through its algorithms to send the connection to a specific application instance on a server. The load balancer continues to manage and monitor the connection for the entire duration.

Load balancers health check the application on the server to determine its availability. If the health check fails, the load balancer takes that instance of the application out of its pool of available servers. When the application comes back online, the health check validates its availability and the server is put back into the availability pool.

This project is tested on locally hosted servers that runs an instance of test-application.

#### Things you should know to test it
- pass "PORT" & "Server-Name" along with node running command and modify the server credentials in Available server's List.
- LoadBalancer runs on PORT ":8000" , you can host it elsewhere as well.
- requests that goes to the LoadBalancer is directly redirected to one of the server's in the pool of available servers.
- In HealthCheck, periodic calls are implemented using a package called 'gocron' you can refer it here(https://github.com/go-co-op/gocron).
  




