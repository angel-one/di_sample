# di_sample

This is Java Style Bean DI sample using just reflect and a helper Beans Repository class.

1. Inidividual Beans ``Register()`` themselves during  individual ``init()`` function
2. In the main.go function, we call one 'Initialize()' function on the beans repo which does DFS on the dependency DAG and does init


