# Single-line comments use a hashtag
// C style comments work too
/*
 * And multi-line comments are just like C.
 */

# Namespaces are language-specific. Generated code uses the namespace
# to separate generated code into packages or modules.
namespace java com.soofaloofa.taskmanager

/**
 * Define some common types for our service.
 * Thrift ensures that different languages serialize these type definitions to appropriate
 * language-specific types.
 */
typedef i64 id
typedef i32 int

/**
 * Define a constant using the `const` keyword.
 */
const id DEFAULT_ID = -1

/**
 * structs are mapped by Thrift to classes or structs in your language of
 * choice. This struct has two fields, an Identifier of type `id` and
 * a Description of type `string`. The Identifier defaults to DEFAULT_ID.
 */
struct Task {
    /** A unique identifier for this task. */
    1: id Identifier = DEFAULT_ID,

    /** A description of the task. */
    2: string Description
}

/**
 * Exceptions inherit from language-specific base exceptions.
 */
exception TaskException {
    /**@ The reason for this exception. */
    1: string Reason
}

/**
 * A service defines the API that is exposed to clients.
 *
 * This TaskManager service has one available endpoint for creating a task.
 */
service TaskManager {
    /**@
     * Create a new task.
     *
     * This method accepts a Task struct and returns an i64.
     * It may throw a TaskException.
     */
    i64 create(1:Task task) throws (1:TaskException e)
}
