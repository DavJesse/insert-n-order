<h1><strong><span style="font-size:23pt;">String Manipulation with Flags</span></strong></h1>
<p><span style="font-size:11pt;">This program is designed to manipulate strings based on command-line flags provided as arguments. It supports the following functionalities:</span></p>
<ul>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="font-size:11pt;">Insertion of a string into another string</span></p>
    </li>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="font-size:11pt;">Ordering of a string in ASCII order</span></p>
    </li>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="font-size:11pt;">Printing help information when no arguments are provided or when the&nbsp;</span><span style="color:#188038;font-size:11pt;">--help</span><span style="font-size:11pt;">&nbsp;flag is used</span></p>
    </li>
</ul>
<h2><strong><span style="font-size:17pt;">Usage</span></strong></h2>
<p><span style="font-size:11pt;">The program can be executed with the following command:</span></p>
<pre>
<span style="font-size:11pt;">bash</span>
<span style="color:#188038;font-size:11pt;">go run . [flags] [string]</span>
</pre>
<h3><strong><span style="font-size:13pt;">Flags</span></strong></h3>
<ul>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="color:#188038;font-size:11pt;">--insert</span><span style="font-size:11pt;">&nbsp;or&nbsp;</span><span style="color:#188038;font-size:11pt;">-i</span><span style="font-size:11pt;">: This flag inserts the string into the string passed as an argument.</span></p>
    </li>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="color:#188038;font-size:11pt;">--order</span><span style="font-size:11pt;">&nbsp;or&nbsp;</span><span style="color:#188038;font-size:11pt;">-o</span><span style="font-size:11pt;">: This flag will order the string argument in ASCII order.</span></p>
    </li>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="color:#188038;font-size:11pt;">--help</span><span style="font-size:11pt;">&nbsp;or&nbsp;</span><span style="color:#188038;font-size:11pt;">-h</span><span style="font-size:11pt;">: Displays the usage instructions along with explanations of the available flags.</span></p>
    </li>
</ul>
<h3><strong><span style="font-size:13pt;">Examples</span></strong></h3>
<pre>
<span style="font-size:11pt;">bash</span>
<span style="color:#188038;font-size:11pt;">$ go run . --insert=4321 --order asdad</span>

<span style="color:#188038;font-size:11pt;">1234aadds</span>

<span style="color:#188038;font-size:11pt;">$ go run . --insert=4321 asdad</span>
<span style="color:#188038;font-size:11pt;">asdad4321</span>

<span style="color:#188038;font-size:11pt;">$ go run . asdad</span>
<span style="color:#188038;font-size:11pt;">asdad</span>

<span style="color:#188038;font-size:11pt;">$ go run . --order 43a21</span>
<span style="color:#188038;font-size:11pt;">1234a</span>

<span style="color:#188038;font-size:11pt;">$ go run . --help</span>
<span style="color:#188038;font-size:11pt;">--insert</span>
<span style="color:#188038;font-size:11pt;">&nbsp; -i</span>
<span style="color:#188038;font-size:11pt;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;This flag inserts the string into the string passed as an argument.</span>
<span style="color:#188038;font-size:11pt;">--order</span>
<span style="color:#188038;font-size:11pt;">&nbsp; -o</span>
<span style="color:#188038;font-size:11pt;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;This flag will behave like a boolean; if it is called, it will order the argument.</span>
</pre>
<h2><strong><span style="font-size:17pt;">Implementation Details</span></strong></h2>
<ul>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="font-size:11pt;">The program processes command-line arguments to determine the action to perform.</span></p>
    </li>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="font-size:11pt;">If no arguments are provided or if the&nbsp;</span><span style="color:#188038;font-size:11pt;">--help</span><span style="font-size:11pt;">&nbsp;flag is used, it prints usage instructions.</span></p>
    </li>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="font-size:11pt;">Multiple flags can be used simultaneously, and the string manipulation is performed based on the order of appearance of flags in the arguments.</span></p>
    </li>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="font-size:11pt;">The insert flag inserts the attached string at the end of the argument string.</span></p>
    </li>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="font-size:11pt;">Use the following format to use the insert flag:</span></p>
    </li>
</ul>
<pre>
<span style="font-size:11pt;">bash</span>
<span style="color:#188038;font-size:11pt;">--insert=[string]</span>
<span style="color:#188038;font-size:11pt;">-i=[string]</span>
</pre>
<p><span style="font-size:11pt;">Where&nbsp;</span><span style="color:#188038;font-size:11pt;">[string]&nbsp;</span><span style="font-size:11pt;">is the string you wish to to attach.</span></p>
<h2><strong><span style="font-size:17pt;">Customization</span></strong></h2>
<p><span style="font-size:11pt;">Feel free to customize the program according to your needs. You can modify the code to add new functionalities or improve existing ones. Additionally, you can extend the program to support more command-line flags or enhance the string manipulation logic.</span></p>
<h2><strong><span style="font-size:17pt;">License</span></strong></h2>
<p><span style="font-size:11pt;">This program is open-source and available under the MIT License. Feel free to use, modify, and distribute it for any purpose. Contributions and feedback are welcome!</span></p>
<h2><strong><span style="font-size:17pt;">Author</span></strong></h2>
<span style="font-size:11pt;">By David Jesse Odhiambo</span></p>
<p><span style="font-size:11pt;">Apprentice Software Devoloper Zone01 Kisumu</span></p>

<hr>
<p><br></p>