package server

// Owner: Member B
//
// Welcome text, Linux logo, and name prompt (copied exactly from the subject).

// Logo contains the Linux penguin (Tux) ASCII art.
const Logo = "         _nnnn_\n" +
	"        dGGGGMMb\n" +
	"       @p~qp~~qMb\n" +
	"       M|@||@) M|\n" +
	"       @,----.JM|\n" +
	"      JS^\\__/  qKL\n" +
	"     dZP        qKRb\n" +
	"    dZP          qKKb\n" +
	"   fZP            SMMb\n" +
	"   HZM            MMMM\n" +
	"   FqM            MMMM\n" +
	" __| \".        |\\dS\"qML\n" +
	" |    `.       | `' \\Zq\n" +
	"_)      \\.___.,|     .'\n" +
	"\\____   )MMMMMP|   .'\n" +
	"     `-'       `--'\n"

// WelcomeMessage is sent to a client right after connecting.
const WelcomeMessage = "Welcome to TCP-Chat!\n" + Logo

// PromptName asks the connecting client to enter a username.
const PromptName = "[ENTER YOUR NAME]: "
