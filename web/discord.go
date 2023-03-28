func main() {
    // send all logs to stdout
    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

    discord_client_secret = '8dyfuiRyq=vVc3RRr_edRk-fK__JItpZ' 

    cmd.Execute()
}
