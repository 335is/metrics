// test.cpp : Defines the entry point for the application.
//

#include <iostream>
#include <chrono>
#include <thread>

#include "StatsdClient.hpp"
using namespace Statsd;

using namespace std;

int main(int argc, char* argv[])
{
	cout << "C++ cpp-client metrics client" << "\n";

	std::string ip("127.0.0.1");
	uint16_t port(8125);
	std::string prefix("cpp-client");
	std::string metricName("main.forloop.increment");
	StatsdClient client{ ip, port, prefix };

	chrono::milliseconds wait_time(1000);

	while (true)
	{
		client.increment(metricName);
		cout << "Incrementing metric " << prefix << "." << metricName << ":" << "1 => StatsD at " << ip << ":" << port << "\n";
		this_thread::sleep_for(wait_time);
	}

	return 0;
}
