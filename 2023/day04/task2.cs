using System;
using System.IO;
using System.Collections.Generic;

class Task1
{
    static void Main(string[] args)
    {
        List<string> cards = GetData(args);

        int sum = 0;
        int[] nrOfCards = new int[cards.Count];

        for (int i = 0; i < cards.Count; i++)
        {
            nrOfCards[i] = 1;
        }

        for (int i = 0; i < cards.Count; i++)
        {
            int cardScore = 0;
            string line = cards[i];
            List<string> nStr = new List<string>();
            HashSet<string> winningNumbers = new HashSet<string>();

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
                    cardScore += 1;
                }
            }

            for (int j = i + 1; j <= i + cardScore; j++)
            {
                nrOfCards[j] += 1 * nrOfCards[i];
            }
        }

        foreach (int n in nrOfCards)
        {
            sum += n;
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
