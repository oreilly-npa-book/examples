# Python contains very useful tools for working with JSON, and they're
# part of the standard library, meaning they're built into Python itself.
import json

# We can load our JSON file into a variable called "data"
with open("data.json") as f:
    data = f.read()

# json_dict is a dictionary, and json.loads takes care of
# placing our JSON data into it.
json_dict = json.loads(data)

# Printing information about the resulting Python data structure
print("The entire JSON document is loaded as type %s\n" % type(json_dict))
print ("Now printing each item in this document and the type it contains")
for k, v in json_dict.items():
    print("-- The key %s contains a %s value." % (str(k), str(type(v))))
