<h1><a id="Tribonachi_0"></a>Tribonachi</h1>
<p>Can calculate large order of the tribonacci number<br>
Really big (up to 109999 characters)</p>
<h1><a id="Features_5"></a>Features!</h1>
<ul>
<li>Simple start. where man n = 75 (limitation uint64);</li>
<li>Configure with a simple utility;</li>
<li>Running in the docker or docker-compose;</li>
</ul>
<h1><a id="Get_started_11"></a>Get started</h1>
<ol>
<li>Сopy this repository</li>
<li>Install the necessary libraries</li>
<li>Build project</li>
<li>Configure with</li>
</ol>
<pre><code class="language-sh">$ git <span class="hljs-built_in">clone</span> git@github.com:Lasiar/tribonachi.git
$ <span class="hljs-built_in">cd</span> tribonachi
$ make
$ <span class="hljs-built_in">cd</span> ./bin
$ ./configure <span class="hljs-comment">#answer the questions</span>
$ sudo docker-compose up --build
</code></pre>
<p>optional for redis:</p>
<pre><code># echo never &gt; /sys/kernel/mm/transparent_hugepage/enabled
</code></pre>
<p>Open the browser, in the address bar, type 127.0.0.1, and the port you specified is the default 8080.<br>
Using the get query, the value n (exponent)</p>
<p>example:</p>
<pre><code>127.0.0.1/tribonachi?n=10
</code></pre>
<h3><a id="Requirement_38"></a>Requirement</h3>
<ul>
<li>Golang 1.9.4</li>
<li>Docker 17.12.1-ce (optional)</li>
<li>docker-compose version 1.19.0, build 9e633ef (optional)</li>
<li>docker-py version: 2.7.0 (optional)</li>
<li>CPython version: 2.7.13 (optional)</li>
<li>OpenSSL version: OpenSSL 1.0.1t  3 May 2016 (optional)</li>
</ul>
<p>Warning: demanding memory, at least 16 gigabytes</p>
<h3><a id="Todos_48"></a>Todos</h3>
<ul>
<li>Fix volumes</li>
<li>Add limitation on the memory of the redis (usage with info), relative to the total</li>
<li>Refactor http.web.HttpServer (looks ugly)</li>
<li>Add validation in configure</li>
</ul>
<h3><a id="Known_issue_55"></a>Known issue</h3>
<ul>
<li>Dont work volumes redis</li>
<li>Recalculation each launch</li>
</ul>
<h2><a id="License_59"></a>License</h2>
<p>MIT</p>
<p><strong>Free Software</strong></p>
