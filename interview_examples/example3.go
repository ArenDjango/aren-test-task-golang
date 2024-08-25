package main

/*
	127.0.0.1
	127.0.0.2
	127.0.0.3
*/

// concurrently get simultanisly from 3 replicas and if one of this answering, stop remainings

//const timeout = time.Second
//
//var ErrNotFound = errors.New("not found")
//var ErrDistributedGetFailed = errors.New("ErrDistributedGetFailed")
//
//func Get(ctx context.Context, address, key string) (string, error) {
//	time.Sleep(2 * time.Second)
//	return "", nil
//}
//
//func DistributiveGet(ctx context.Context, addresses []string, key string) (string, error) {
//	if ctx.Err() != nil {
//		return "", ctx.Err()
//	}
//
//	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
//	defer cancel()
//
//	var wg sync.WaitGroup
//
//	wg.Add(len(addresses))
//	valueCh := make(chan string)
//	notFoundErrCh := make(chan error)
//
//	go func() {
//		wg.Wait()
//		close(valueCh)
//		close(notFoundErrCh)
//	}()
//
//	for _, address := range addresses {
//		go func(address string) {
//			defer wg.Done()
//			value, err := Get(ctxWithTimeout, address, key)
//			if err != nil && errors.Is(err, ErrNotFound) {
//				notFoundErrCh <- err
//				return
//			}
//			select {
//			case valueCh <- value:
//				close(valueCh)
//			default:
//			}
//
//		}(address)
//	}
//
//	select {
//	case <-ctxWithTimeout.Done():
//		return "", ctxWithTimeout.Err()
//	case v := <-valueCh:
//		return v, nil
//	case err := <-notFoundErrCh:
//		return "", err
//	}
//}
//
//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//
//	v, err := DistributiveGet(ctx, []string{"127.0.0.1", "127.0.0.2", "127.0.0.3"}, "posts")
//	if err != nil {
//		fmt.Errorf("err to get something")
//	}
//	fmt.Println(v)
//}
