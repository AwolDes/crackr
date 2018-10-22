for cpu in 1 2 3 4 5 6 7 8
do  
    echo "=================="
    echo "=== $cpu PROCS ==="
    echo "=================="
    for i in 1 2 3 4 5
    do
        time GOMAXPROCS=$cpu crackr -hf hashes.txt -d dictionaries/
    done
done
