# import this


print("Hello world!!")


def introduce(name:str, languages:list[str]) -> str:
    """
    Returns a formatted introduction string.
    Note: tripe quotes = docstring (built-in documentation)
    """
    lang_str = ", ".join(languages)
    return f"Hi, I'm {name}. I know: {lang_str}"

if __name__ == "__main__":
    message = introduce("Arjun", ["C++", "Go", "Python"])
    print(message)