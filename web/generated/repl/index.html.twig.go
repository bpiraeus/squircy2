// Code generated by stickgen.
// DO NOT EDIT!

package repl

import (
	"fmt"
	"io"

	"github.com/tyler-sommer/stick"
)

func blockReplIndexHtmlTwigContent(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
	// line 3, offset 19 in repl/index.html.twig
	fmt.Fprint(output, `
<div class="row">
	<div class="col-md-8" style="height: 300px">
		<h4>REPL</h4>
		<select class="form-control" id="script_type" name="scriptType">
			<option>Javascript</option>
		</select>
		<br>
  		<div id="editor" style="width: 100%; height: 100%"></div>
        <div style="display: none">
            <textarea id="code-body"></textarea>
        </div>
		<br>
		<a class="form-control btn btn-primary" id="execute" href="/repl/execute">Execute</a>
	</div>
	<div class="col-md-4">
		<h5>Output</h5>
		<pre id="output">
		</pre>
		<h5>Events</h5>
		<pre id="event-log" class="history"></pre>
	</div>
</div>
`)
}
func blockReplIndexHtmlTwigAdditionalJavascripts(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
	// line 28, offset 34 in repl/index.html.twig
	fmt.Fprint(output, `
<script src="//cdn.jsdelivr.net/ace/1.1.7/min/ace.js"></script>
<script src="//cdn.jsdelivr.net/ace/1.1.7/min/ext-searchbox.js"></script>
<script src="//cdn.jsdelivr.net/ace/1.1.7/min/ext-spellcheck.js"></script>
<script src="//cdn.jsdelivr.net/ace/1.1.7/min/ext-static_highlight.js"></script>
<script src="//cdn.jsdelivr.net/ace/1.1.7/min/mode-javascript.js"></script>
<script src="//cdn.jsdelivr.net/ace/1.1.7/min/theme-textmate.js"></script>
<script src="//cdn.jsdelivr.net/ace/1.1.7/min/worker-javascript.js"></script>
<script type="text/javascript">
$(function() {
	var modeMap = {
		"Javascript": "ace/mode/javascript",
	};
	var $typeField = $('#script_type');
	var $bodyField = $('#code-body');
	var editor = ace.edit("editor");
    editor.setTheme("ace/theme/textmate");
	editor.resize();
	editor.setValue($bodyField.val());
	editor.getSession().on('change', function() {
		$bodyField.val(editor.getValue());
	});
	
	$typeField.on('change', function() {
		editor.getSession().setMode(modeMap[$typeField.val()]);
	}).change();
	
	var $execute  = $('#execute');
	var $output   = $('#output');
	var $eventLog = $('#event-log');

	var es = window.squIRCyEvents;

	es.addEventListener("irc.WILDCARD", function(e) {
		var data = JSON.parse(e.data);
		$eventLog.append("[" + data.Code + "] " + data.Nick + "->" + data.Target + ": " + data.Message + "\n");
		$eventLog[0].scrollTop = $eventLog[0].scrollHeight;
	});
	
	$execute.on('click', function(e) {
		e.preventDefault();
		
		$.ajax({
			url: $execute.attr('href'),
			type: 'post',
			data: {
				script: editor.getValue(),
				scriptType: $typeField.val()
			},
			success: function(response) {
				$output.html(JSON.stringify(response, null, '  '))
			}
		});
	});
});
</script>
`)
}

func TemplateReplIndexHtmlTwig(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
	// line 1, offset 0 in layout.html.twig
	fmt.Fprint(output, `<!DOCTYPE html>
<html>
<head>
  <title>squIRCy</title>
  <script src="//cdn.jsdelivr.net/jquery/2.1.1/jquery.min.js"></script>
  <script src="//cdn.jsdelivr.net/bootstrap/3.2.0/js/bootstrap.min.js"></script>
  <script src="//cdn.jsdelivr.net/momentjs/2.8.1/moment.min.js"></script>
  <link rel="stylesheet" href="//cdn.jsdelivr.net/bootstrap/3.2.0/css/bootstrap.min.css">
  <link rel="stylesheet" href="//cdn.jsdelivr.net/fontawesome/4.2.0/css/font-awesome.min.css">
  <link rel="stylesheet" href="/css/style.css">
</head>

<body>
	<div id="main-container" class="container">
		<div class="row">
			<div class="col-md-12">
			`)
	// line 17, offset 6 in layout.html.twig
	blockReplIndexHtmlTwigContent(env, output, ctx)
	// line 18, offset 17 in layout.html.twig
	fmt.Fprint(output, `
			</div>
		</div>
	</div>

	<nav id="main-nav" class="navbar navbar-default navbar-fixed-top" role="navigation">
	  	<div class="container">
			<div class="navbar-header">
				<a class="navbar-brand" href="https://github.com/tyler-sommer/squircy2">squIRCy2</a>
        	</div>
			<ul class="nav navbar-nav">
				<li><a href="/">Dashboard</a></li>
				<li><a href="/script">Scripts</a></li>
				<li><a href="/webhook">Webhooks</a></li>
				<li><a href="/manage">Settings</a></li>
				<li class="divider">&nbsp;</li>
				<li><a href="/repl">REPL</a></li>
			</ul>
			<ul class="nav navbar-nav pull-right">
				<li><a id="reinit" title="Re-initialize scripts" href="/script/reinit"><i class="fa fa-refresh"></i></a></li>
		        <li><a class="post-action" id="connect-button" style="display: none" href="/connect"><i class="fa fa-power-off"></i> Connect</a></li>
				<li><a class="post-action" id="disconnect-button" style="display: none" href="/disconnect"><i class="fa fa-power-off"></i> Disconnect</a></li>
				<li><a class="post-action" id="connecting-button" style="display: none" href="/disconnect"><i class="fa fa-power-off"></i> Connecting...</a></li>
		      </ul>
	  	</div>
	</nav>

    `)
	// line 1, offset 0 in _page_javascripts.html.twig
	fmt.Fprint(output, `<script type="text/javascript">
$(function() {
    var es = window.squIRCyEvents = new EventSource('/event');

	var $connectBtn = $('#connect-button');
	var $disconnectBtn = $('#disconnect-button');
	var $connectingBtn = $('#connecting-button');
    es.addEventListener("irc.CONNECTING", function() {
        $connectingBtn.show();
		$disconnectBtn.add($connectBtn).hide();
    });
    es.addEventListener("irc.CONNECT", function() {
        $disconnectBtn.show();
		$connectBtn.add($connectingBtn).hide();
    });
    es.addEventListener("irc.DISCONNECT", function() {
        $connectBtn.show();
        $disconnectBtn.add($connectingBtn).hide();
    });

	var $postLinks = $('.post-action');
	$postLinks.on('click', function(e) {
		e.preventDefault();
		
		var url = $(this).attr('href');
		$.ajax({
			url: url,
			type: 'post'
		});
	});
	
	var $reinit = $('#reinit').tooltip({
		placement: 'bottom',
		container: 'body'
	});
	$reinit.on('click', function(e) {
		e.preventDefault();
		
		if (confirm('Are you sure you want to reinitialize all scripts?')) {
			var url = $(this).attr('href');
			$.ajax({
				url: url,
				type: 'post'
			})
		}
	});

	$.ajax({
		url: '/status',
		type: 'get',
		dataType: 'json',
		success: function(response) {
			if (response.Status == 2) {
				$disconnectBtn.show();
				$connectBtn.add($connectingBtn).hide();
			} else if (response.Status == 1) {
				$connectingBtn.show();
				$disconnectBtn.add($connectBtn).hide();
			} else {
				$connectBtn.show();
				$disconnectBtn.add($connectingBtn).hide();
			}
		}
	});
});
</script>`)
	// line 45, offset 47 in layout.html.twig
	fmt.Fprint(output, `
    `)
	// line 46, offset 7 in layout.html.twig
	blockReplIndexHtmlTwigAdditionalJavascripts(env, output, ctx)
	// line 47, offset 18 in layout.html.twig
	fmt.Fprint(output, `
</body>

</html>
`)
	// line 1, offset 32 in repl/index.html.twig
	fmt.Fprint(output, `

`)
	// line 26, offset 14 in repl/index.html.twig
	fmt.Fprint(output, `

`)
	// line 84, offset 14 in repl/index.html.twig
	fmt.Fprint(output, `
`)
}
