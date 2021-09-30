using System;
using System.Threading;
using JustEat.StatsD;

namespace JustEat
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("C# JustEat metrics client");

            var statsDConfig = new StatsDConfiguration { Host = "127.0.0.1", Port = 8125, Prefix = "justeat-client" };
            var address = statsDConfig.Host + ":" + statsDConfig.Port;
            string metricName = "main.forloop.increment";
            IStatsDPublisher statsDPublisher = new StatsDPublisher(statsDConfig);

            for (long metricValue = 0; true; ++metricValue)
            {
                IStatsDPublisherExtensions.Increment(statsDPublisher, metricValue, metricName);
                Console.WriteLine("Incrementing metric {0}.{1}:{2} => StatsD at {3}", statsDConfig.Prefix, metricName, metricValue, address);
                Thread.Sleep(1000);
            }
        }
    }
}
