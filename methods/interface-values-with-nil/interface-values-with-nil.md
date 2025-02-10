# Interface values with nil underlying values

In some languages, a null pointer exception is thrown when concrete value within the interface is nil

In Go, is common creating methods that gracefully handle being called with a nil receiver
