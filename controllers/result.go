package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

// Requeue triggers a object requeue.
func Requeue() (ctrl.Result, error) {
	return ctrl.Result{Requeue: true}, nil
}

// RequeueWithError triggers a object requeue because the informed error happend.
func RequeueWithError(err error) (ctrl.Result, error) {
	return ctrl.Result{Requeue: true}, err
}

// NoRequeue all done, the object does not need reconciliation anymore.
func NoRequeue() (ctrl.Result, error) {
	return ctrl.Result{Requeue: false}, nil
}

// NoRequeueWithError does not trigger a requeue, but an error is informed.
func NoRequeueWithError(err error) (ctrl.Result, error) {
	return ctrl.Result{Requeue: false}, err
}
