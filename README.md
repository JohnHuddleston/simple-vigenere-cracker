# simple-vigenere-cracker
A simple Go program that generates and tests all possible rotation keys for a Vigenere cipher using a key of length 5.  Originally used to be lazy about a university assignment.  This particular program is extremely inelegant and exploits a specific property of the given executable (the output of a success message given the correct key).  Without this your best bet would be to do some level of language processing or to use some kind of English confidence metric, which may increase complexity beyond a reasonable level.

This program really served 2 purposes:
* Let me be lazy and not have to count letters to do frequency analysis on columnnated ciphertext, and
* provide a project for me to really start learning Go

Given that my back of the envelope calculations put the maximum runtime on my laptop at around 5 hours (a similar Python script ran for 12 to no avail) and ample time was provided for the assignment, this absolutely succeeded on the first purpose.  ~~**But**, I'm not satisfied with how much I learned, primarily in terms of multithreading.  I kept running into type errors while attempting to partition beyond manually calculated values (due to float64s in the mix because of the large number of keys), so that's currently on my to-do list.  Go routines are amazingly scalable and I'd like to spin up hundreds of them to at least find a point of diminishing returns to optimize this as much as it can be.

The script has since been updated to dynamically partition the set of all possible keys based on a variable that you can change to reflect how many cores you have to work with.  If I really get the urge to I'll probably modify the program to utilize some flavor of MPI in order to allow computation across multiple nodes.  That said, I do not currently have access to a cluster capable of running the assignment's executable (damn libgcrypt!) so testing will take some creative problem solving.

**Note**: In most cases it's probably a lot more effective to just create sets of ciphertext characters based on the length of the key then just look at the occurences.  In most cases you can just map the most frequent letter to E and know the rotation of that column, rinse and repeat for the rest.  If you're given enough text to work with this should be fine and quick.
