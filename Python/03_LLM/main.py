from transformers import AutoModelForCausalLM, AutoTokenizer
# Load model and tokenizer
model = AutoModelForCausalLM.from_pretrained(
"microsoft/Phi-3-mini-4k-instruct",
device_map="auto",
torch_dtype="auto",
trust_remote_code=True,
)
tokenizer = AutoTokenizer.from_pretrained("microsoft/Phi-3-mini-4k-instruct")


prompt = "Write an email apologizing to Sarah for the tragic gardening mishap. Explain how it happened.<|assistant|>"
# Tokenize the input prompt
input_ids = tokenizer(prompt, return_tensors="pt").input_ids.to(model.device)
# Generate the text
generation_output = model.generate(
input_ids=input_ids,
max_new_tokens=20
)
# Print the output
print(tokenizer.decode(generation_output[0]))