import csv
 
# open the file in read mode
filename = open('dump.txt', 'r')
 
# creating dictreader object
file = csv.DictReader(filename)
 
# creating empty lists
out = ""
values = []
 
# iterating over each row and append
# values to empty list
for col in file:
    values.append(col['ASCII'])

# printing lists
print(out.join(values))