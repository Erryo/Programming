import random


def get_guess():
    while True:
        guess = input("Input a 5 letter word: ")
        if len(guess) != 5:
            print("Length has to be EQUAL to 5!!!")
            continue
        if not guess.isalpha():
            print("Do NOT use numbers!!!")
            continue
        break
    return guess


def check_colors(guess, word):
    if guess == word:
        return [[], [0, 1, 2, 3, 4]]
    characters = []
    colors = [[], []]  # [[yelow index],[green index]]
    for char_ind in range(len(guess)):
        for letter_ind in range(len(word)):
            char = guess[char_ind]
            letter = word[letter_ind]
            if char in characters:
                continue
            if char != letter:
                continue
            if char_ind != letter_ind:
                print(f"Char:{char} of {guess} is same as letter {letter} of {word} ")
                colors[0].append(char_ind)
                characters.append(char)
                continue
            print(f"Green at char {char}, index {char_ind} ")
            colors[1].append(char_ind)
    return colors


def draw_mat(guess, colors):
    print(guess)
    for index in range(len(guess)):
        if index in colors[0]:
            print(guess[index], " : ", "y")
            continue
        if index in colors[1]:
            print(guess[index], " : ", "g")
            continue
        print(guess[index], " : ", "-")


dictionary = ["carte", "vreme", "salut"]
# word = dictionary[random.randint(0, len(dictionary) - 1)]
word = "dweeb"
print(word)
while True:
    user_guess = get_guess()
    colors = check_colors(user_guess, word)
    draw_mat(user_guess, colors)
    if colors[1] == [0, 1, 2, 3, 4]:
        print("You have won")
        break
