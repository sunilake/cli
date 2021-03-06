// This file was generated by counterfeiter
package fake_authenticators

import (
	"regexp"
	"sync"

	"code.cloudfoundry.org/diego-ssh/authenticators"
	"golang.org/x/crypto/ssh"
)

type FakePasswordAuthenticator struct {
	UserRegexpStub        func() *regexp.Regexp
	userRegexpMutex       sync.RWMutex
	userRegexpArgsForCall []struct{}
	userRegexpReturns     struct {
		result1 *regexp.Regexp
	}
	AuthenticateStub        func(metadata ssh.ConnMetadata, password []byte) (*ssh.Permissions, error)
	authenticateMutex       sync.RWMutex
	authenticateArgsForCall []struct {
		metadata ssh.ConnMetadata
		password []byte
	}
	authenticateReturns struct {
		result1 *ssh.Permissions
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePasswordAuthenticator) UserRegexp() *regexp.Regexp {
	fake.userRegexpMutex.Lock()
	fake.userRegexpArgsForCall = append(fake.userRegexpArgsForCall, struct{}{})
	fake.recordInvocation("UserRegexp", []interface{}{})
	fake.userRegexpMutex.Unlock()
	if fake.UserRegexpStub != nil {
		return fake.UserRegexpStub()
	} else {
		return fake.userRegexpReturns.result1
	}
}

func (fake *FakePasswordAuthenticator) UserRegexpCallCount() int {
	fake.userRegexpMutex.RLock()
	defer fake.userRegexpMutex.RUnlock()
	return len(fake.userRegexpArgsForCall)
}

func (fake *FakePasswordAuthenticator) UserRegexpReturns(result1 *regexp.Regexp) {
	fake.UserRegexpStub = nil
	fake.userRegexpReturns = struct {
		result1 *regexp.Regexp
	}{result1}
}

func (fake *FakePasswordAuthenticator) Authenticate(metadata ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	var passwordCopy []byte
	if password != nil {
		passwordCopy = make([]byte, len(password))
		copy(passwordCopy, password)
	}
	fake.authenticateMutex.Lock()
	fake.authenticateArgsForCall = append(fake.authenticateArgsForCall, struct {
		metadata ssh.ConnMetadata
		password []byte
	}{metadata, passwordCopy})
	fake.recordInvocation("Authenticate", []interface{}{metadata, passwordCopy})
	fake.authenticateMutex.Unlock()
	if fake.AuthenticateStub != nil {
		return fake.AuthenticateStub(metadata, password)
	} else {
		return fake.authenticateReturns.result1, fake.authenticateReturns.result2
	}
}

func (fake *FakePasswordAuthenticator) AuthenticateCallCount() int {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return len(fake.authenticateArgsForCall)
}

func (fake *FakePasswordAuthenticator) AuthenticateArgsForCall(i int) (ssh.ConnMetadata, []byte) {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return fake.authenticateArgsForCall[i].metadata, fake.authenticateArgsForCall[i].password
}

func (fake *FakePasswordAuthenticator) AuthenticateReturns(result1 *ssh.Permissions, result2 error) {
	fake.AuthenticateStub = nil
	fake.authenticateReturns = struct {
		result1 *ssh.Permissions
		result2 error
	}{result1, result2}
}

func (fake *FakePasswordAuthenticator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.userRegexpMutex.RLock()
	defer fake.userRegexpMutex.RUnlock()
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePasswordAuthenticator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ authenticators.PasswordAuthenticator = new(FakePasswordAuthenticator)
