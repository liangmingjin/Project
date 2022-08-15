Dim Program1 Program1="E:\聊天工具\QQ2009\Bin\QQ.exe" 
Set WshShell=createobject("wscript.shell") 
Set oExec=WshShell.Exec(Program1) 
WScript.Sleep 2000 
WshShell.AppActivate "Q登录" 
WshShell.SendKeys "123456" WScript.Sleep 200 WshShell.SendKeys "{TAB}" WshShell.SendKeys "123456" WScript.Sleep 200 WshShell.SendKeys "{TAB}" WshShell.SendKeys "{Enter}" WshShell.SendKeys "{DOWN}" WshShell.SendKeys "{DOWN}" WshShell.SendKeys "{DOWN}" WshShell.SendKeys "{DOWN}" WshShell.SendKeys "{DOWN}" WshShell.SendKeys "{DOWN}" WshShell.SendKeys "{Enter}" WScript.Sleep 200 WshShell.SendKeys "{TAB}" 