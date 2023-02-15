import numpy as np


def play_game(numbers: list, cards: list) -> list:
    print("Playing game ...")
    results = []

    for num in numbers:

        for card in cards:

            # Only consider card that have not already called bingo
            if not card["bingo"]:

                # If the number exists replace with 100
                card["card"] = np.where(card["card"] == num, 100, card["card"])

                # Check the row and column totals
                # If 500 then we have bingo
                # Calculate the result and set card to bingo
                sum_rows = np.sum(card["card"], axis=1)
                sum_cols = np.sum(card["card"], axis=0)

                if 500 in sum_rows or 500 in sum_cols:

                    # Set all the 100 values to 0
                    card["card"] = np.where(card["card"] == 100, 0, card["card"])

                    # Get the sum of the card
                    sum_card = np.sum(card["card"])

                    results.append(sum_card * num)
                    card["bingo"] = True

    return results


if __name__ == "__main__":

    with open("input") as f:

        raw_data = f.read().split("\n\n")

        numbers = raw_data[0]
        list_of_numbers = [int(n) for n in numbers.split(",")]

        list_of_cards = []

        for card in raw_data[1:]:

            rows = card.split("\n")
            card_array = np.array([[int(n) for n in (row.split())] for row in rows])
            card = {"card": card_array, "bingo": False}
            list_of_cards.append(card)

    results = play_game(list_of_numbers, list_of_cards)

    print(f"PART 01: {results[0]}")
    print(f"PART 02: {results[-1]}")