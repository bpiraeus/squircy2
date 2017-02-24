// Code generated by stickgen.
// DO NOT EDIT!

package generated

import (
	"fmt"
	"io"

	"github.com/tyler-sommer/stick"
)

func blockIndexHtmlTwigContent(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
	// line 3, offset 19 in index.html.twig
	fmt.Fprint(output, `
<div class="row">
	<div class="col-md-8">
		<h4>Dashboard</h4>
	</div>
</div>
<div class="row">
    <div class="col-md-12">
        <h5>Terminal</h5>
        <pre id="terminal-log" class="history">`)
	// line 12, offset 50 in index.html.twig
	stick.Iterate(ctx["terminal"], func(_, line stick.Value, loop stick.Loop) (brk bool, err error) {
		// line 12, offset 73 in index.html.twig
		{
			val, err := stick.GetAttr(line, "Message")
			if err == nil {
				fmt.Fprint(output, val)
			}
		}
		return false, nil
	})
	// line 12, offset 103 in index.html.twig
	fmt.Fprint(output, `</pre>
        <h5>Events</h5>
        <pre id="event-log" class="history">`)
	// line 14, offset 47 in index.html.twig
	stick.Iterate(ctx["irc"], func(_, line stick.Value, loop stick.Loop) (brk bool, err error) {
		// line 14, offset 65 in index.html.twig
		fmt.Fprint(output, `[`)
		// line 14, offset 66 in index.html.twig
		{
			val, err := stick.GetAttr(line, "Code")
			if err == nil {
				fmt.Fprint(output, val)
			}
		}
		// line 14, offset 81 in index.html.twig
		fmt.Fprint(output, `] `)
		// line 14, offset 83 in index.html.twig
		{
			val, err := stick.GetAttr(line, "Nick")
			if err == nil {
				fmt.Fprint(output, val)
			}
		}
		// line 14, offset 98 in index.html.twig
		fmt.Fprint(output, `->`)
		// line 14, offset 100 in index.html.twig
		{
			val, err := stick.GetAttr(line, "Target")
			if err == nil {
				fmt.Fprint(output, val)
			}
		}
		// line 14, offset 117 in index.html.twig
		fmt.Fprint(output, `: `)
		// line 14, offset 119 in index.html.twig
		{
			val, err := stick.GetAttr(line, "Message")
			if err == nil {
				fmt.Fprint(output, val)
			}
		}
		// line 14, offset 137 in index.html.twig
		fmt.Fprint(output, `
`)
		return false, nil
	})
	// line 15, offset 12 in index.html.twig
	fmt.Fprint(output, `</pre>
    </div>
</div>
</div>
`)
}
func blockIndexHtmlTwigAdditionalJavascripts(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
	// line 21, offset 34 in index.html.twig
	fmt.Fprint(output, `
<script type="text/javascript">
$(function() {
    var $eventLog = $('#event-log');
    var $terminalLog = $('#terminal-log');
    var es = window.squIRCyEvents;
    es.addEventListener("irc.WILDCARD", function(e) {
        var data = JSON.parse(e.data);
        $eventLog.append("[" + data.Code + "] " + data.Nick + "->" + data.Target + ": " + data.Message + "\n");
        $eventLog[0].scrollTop = $eventLog[0].scrollHeight;
    });
    es.addEventListener("cli.OUTPUT", function(e) {
        var data = JSON.parse(e.data);
        $terminalLog.append(data.Message);
        $terminalLog[0].scrollTop = $terminalLog[0].scrollHeight;
    });

    $eventLog[0].scrollTop = $eventLog[0].scrollHeight;
    $terminalLog[0].scrollTop = $terminalLog[0].scrollHeight;
});
</script>
`)
}

func TemplateIndexHtmlTwig(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
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
	blockIndexHtmlTwigContent(env, output, ctx)
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
	blockIndexHtmlTwigAdditionalJavascripts(env, output, ctx)
	// line 47, offset 18 in layout.html.twig
	fmt.Fprint(output, `
</body>

</html>
`)
	// line 1, offset 32 in index.html.twig
	fmt.Fprint(output, `

`)
	// line 19, offset 14 in index.html.twig
	fmt.Fprint(output, `

`)
	// line 42, offset 14 in index.html.twig
	fmt.Fprint(output, `
`)
}