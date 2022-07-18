# Demonstrate the Use of Mutexes in GoLang
Stingy & Spendy are two go routines that share a variable "money". In this
scenario they are racing to increase and decrease the amount of money in the
"money" variable by 10 each time.

Using a locking Mutex we can lock changes to the "money variable" while each of
the methods updated the amount. This way we eliminate the race condition and are
guaranteed that the correct amount of money is add/subtracted once at a time.
