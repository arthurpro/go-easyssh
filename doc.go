/*
Package easyssh provides a simple wrapper around the standard SSH library. Designed to be like net/http but for ssh.
Both server and client implementations are provided.

Creating a client is similar to creating a normal ssh client

  client, err := easyssh.Dial("tcp", "localhost:2022", config)
  if err != nil {
    // Handle error
  }
  defer client.Close()

Once a client is created, you can do a number of things with it:
Local Port Forwarding

  err = client.LocalForward("localhost:8000", "localhost:6060")
  if err != nil {
    // Handle error
  }

Remote Port Forwarding

  err = client.RemoteForward("localhost:8000", "localhost:6060")
  if err != nil {
    // Handle error
  }

Create a session  - used for executing remote commands or getting a remote shell

  session, err := client.NewSession()
  if err != nil {
    // Handle error
  }
  out, err := session.Output("whoami")
  if err != nil {
    // Handle error
  }


Getting started with an SSH server is easy with easyssh

  easyssh.HandleChannel(easyssh.SessionRequest, easyssh.SessionHandler())
  easyssh.HandleChannel(easyssh.DirectForwardRequest, easyssh.DirectPortForwardHandler())
  easyssh.HandleRequestFunc(easyssh.RemoteForwardRequest, easyssh.TCPIPForwardRequest)

  easyssh.ListenAndServe(":2022", sshServerConfig, nil)





There are a lot of layers of ssh communication, and easyssh makes it easy to control at the level desired.


*/
package easyssh
