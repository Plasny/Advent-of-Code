using System;
using System.IO;
using System.Collections.Generic;

class Task1
{
    static void Main(string[] args)
    {
        int sum = 0;
        List<string> cards = GetData(args);

        foreach (string line in cards)
        {
            List<string> nStr = new List<string>();
            HashSet<string> winningNumbers = new HashSet<string>();
            int cardScore = 0;

            nStr.AddRange(line.Split(": ")[1].Split("|"));

            foreach (string n in nStr[0].Split(" "))
            {
                if (n == "") continue;
                winningNumbers.Add(n);
            }

            foreach (string n in nStr[1].Split(" "))
            {
                if (n == "") continue;

                if (winningNumbers.Contains(n))
                {
                    if (cardScore == 0)
                    {
                        cardScore = 1;
                    }
                    else
                    {
                        cardScore *= 2;
                    }
                }
            }

            sum += cardScore;
        }

        Console.WriteLine(sum);
    }

    static List<string> GetData(string[] args)
    {
        List<string> cards = new List<string>();

        if (args.Length == 1)
        {
            try
            {
                cards.AddRange(File.ReadLines(args[0]));
            }
            catch
            {
                Console.WriteLine("Could not open file");
                Environment.Exit(1);
            }
        }
        else
        {
            string line;
            while((line = Console.ReadLine()) != null)
            {
                cards.Add(line);
            }
        }

        return cards;
    }
}
